package helpers

const (
	Duration_Second_1  = "1s"
	Duration_Second_5  = "5s"
	Duration_Second_10 = "10s"
	Duration_Second_30 = "30s"
	Duration_Minute_1  = "1m"
	Duration_Minute_2  = "2m"
	Duration_Minute_5  = "5m"
	Duration_Minute_10 = "10m"
	Duration_Minute_20 = "20m"
	Duration_Hour_1    = "1h"
	Duration_Hour_2    = "2h"
	Duration_Hour_6    = "6h"
	Duration_Hour_24   = "24h"
)

// GetIntervalFromDurationStr returns number of seconds
// in provided duration string
func GetIntervalFromDurationStr(d string) int {
	switch d {
	case Duration_Second_1:
		return 1
	case Duration_Second_5:
		return 5
	case Duration_Second_10:
		return 10
	case Duration_Second_30:
		return 30
	case Duration_Minute_1:
		return 60
	case Duration_Minute_2:
		return 120
	case Duration_Minute_5:
		return 300
	case Duration_Minute_10:
		return 600
	case Duration_Minute_20:
		return 1200
	case Duration_Hour_1:
		return 3600
	case Duration_Hour_2:
		return 7200
	case Duration_Hour_6:
		return 21600
	case Duration_Hour_24:
		return 86400
	default:
		return 0
	}
}