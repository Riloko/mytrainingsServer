package models

import "time"

// Exercise ...
type Exercise struct {
	ID          int
	TrainingID  int
	Image       string
	Name        string
	Description string
	Type        int
	IsGeo       int
	Difficulty  int
	Time        time.Time
}
