package util

import "time"

func Timer(start time.Time, name string) {
	elapsed := time.Since(start).Milliseconds()
	Logger.WithField("timer", name).WithField("starts_on", start.UTC().Format(time.RFC3339)).WithField("elapsed_ms", elapsed).Infoln("Audit")
}
