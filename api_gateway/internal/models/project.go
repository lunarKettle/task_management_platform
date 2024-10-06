package models

import (
	"time"
)

type Project struct {
	Id             uint32
	Name           string
	Description    string
	StartDate      time.Time
	PlannedEndDate time.Time
	ActualEndDate  time.Time
	Status         string
	Priority       uint32
	TeamId         uint32
	Budget         float64
}
