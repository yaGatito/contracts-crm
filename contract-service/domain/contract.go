package domain

import "time"

type Contract struct {
	Number    int64
	StartDate time.Time
	EndDate   time.Time
	Type      string
}
