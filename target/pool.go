package target

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

// Pool holds all the targets.
// This is where the loop operates.
// Key for the Pool is of the format <username>_<targetname>
type Pool struct {
	Mutex   sync.RWMutex
	Targets map[string]Target
}

// NewPool initiates the pool
func NewPool() *Pool {
	return &Pool{
		Targets: make(map[string]Target, 0),
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
func (p *Pool) Get(key string) Target {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	target, targetExists := p.Targets[key]
	if !targetExists {
		return nil
	}
	return target
}

// Adds a new target to the pool
func (p *Pool) Add(tg Target) error {
	if p.Contains(tg.GetPoolKey()) {
		return errors.New("target already exists")
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
	delete(p.Targets, poolKey(userName, targetName))
}

// poolKey returns the key to use in the Pool
// currently it is of the format <userName>_<targetName>
func poolKey(userName, targetName string) string {
	return fmt.Sprintf(
		"%s_%s",
		userName, targetName,
	)
}

// Iterate through all the registered every 10 seconds,
// check if the target needs to be pinged and ping
func (p *Pool) Monitor(ctx context.Context) {
	log.Println("Monitor")
	ticker := time.NewTicker(time.Second * 10)
	targetsChan := make(chan Target, 5000)

	// handle targets in separate goroutine
	go p.handleTargets(ctx, targetsChan)

	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return
		case currentTime := <-ticker.C:
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

func (p *Pool) handleTargets(ctx context.Context, targetsChan <-chan Target) {
	for {
		select {
		case <-ctx.Done():
			return
		case toPingTarget := <-targetsChan:
			go toPingTarget.Ping(ctx)
		}
	}
}
