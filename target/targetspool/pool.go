package targetspool

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"pingmaster/database"
	"pingmaster/target"
)

// Pool holds all the targets.
// This is where the loop operates.
// Key for the Pool is of the format <username>_<targetname>
type Pool struct {
	Mutex   sync.RWMutex
	Targets map[string]target.Target
}

// NewPool initiates the pool
func New() *Pool {
	return &Pool{
		Targets: make(map[string]target.Target, 0),
	}
}

// Contains checks whether the pool contains the target
// for the given userName and targetName
func (p *Pool) Contains(key string) bool {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	target, targetExists := p.Targets[key]
	if !targetExists {
		return false
	}
	if target == nil {
		return false
	}
	return true
}

// Get returns the target for the given userName and targetName
// from the pool if it exists, otherwise it returns nil
func (p *Pool) Get(key string) target.Target {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	target, targetExists := p.Targets[key]
	if !targetExists {
		return nil
	}
	return target
}

// Adds a new target to the pool
func (p *Pool) Add(tg target.Target) error {
	if p.Contains(tg.GetPoolKey()) {
		return errors.New("target with same name already exists")
	}
	p.Mutex.RLock()
	defer p.Mutex.RUnlock()
	p.Targets[tg.GetPoolKey()] = tg
	return nil
}

// Removes the target from the pool
func (p *Pool) Remove(userName, targetName string) {
	p.Mutex.RLock()
	defer p.Mutex.RUnlock()
	delete(p.Targets, target.Key(userName, targetName))
}

// Iterate through all the registered targets every 10 seconds,
// check if the target needs to be pinged and ping
func (p *Pool) Monitor(ctx context.Context, dbConn database.Conn) {
	log.Println("[INF] pool monitor started")

	ticker := time.NewTicker(time.Second * 10)
	targetsChan := make(chan target.Target, 5000)
	pingChan := make(chan *target.Ping, 5000)

	// delete older pings
	go func() {
		deleteTicker := time.NewTicker(time.Minute)
		for range deleteTicker.C {
			err := dbConn.DeleteOldPings(ctx)
			if err != nil {
				log.Printf("[ERR] deleting old pings : %s", err)
			}
		}
	}()

	// ping targets in separate goroutine
	go p.pingTargets(ctx, targetsChan, pingChan)

	// workerpool of 100 workers
	// to insert ping in DB
	for i := 0; i < dbConn.GetMaxConcurrentQueries(); i++ {
		go func() {
			for eachPing := range pingChan {
				err := dbConn.InsertPing(ctx, *eachPing)
				if err != nil {
					log.Printf(
						"[ERR] inserting ping in DB for %s : %s",
						eachPing.TargetKey, err,
					)
				}
			}
		}()
	}

	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return
		case currentTime := <-ticker.C:
			log.Println("[INF] initiating ping loop")
			// retain read lock for the entire loop
			p.Mutex.RLock()
			for _, eachTarget := range p.Targets {
				if eachTarget.NeedToPing(currentTime) {
					// send target to targetsChan to handle separately
					targetsChan <- eachTarget
				}
			}
			p.Mutex.RUnlock()
		}
	}
}

// pingTargets listens for targets to ping on targetsChan
// and once pinged, the same target is then passed into
// dbChan to insert its ping in database
func (p *Pool) pingTargets(ctx context.Context, targetsChan <-chan target.Target, pingChan chan<- *target.Ping) {
	for {
		select {
		case <-ctx.Done():
			log.Println("[INF] stopping ping monitor")
			return
		case toPingTarget := <-targetsChan:
			go func() {
				toPingTarget.Ping(ctx)
				pingChan <- toPingTarget.GetLastPing()
			}()
		}
	}
}
