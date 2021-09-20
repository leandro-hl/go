package lakesite

import "time"

type Job func() error

type JobData struct {
	ID        int
	StartTime time.Time
	EndTime   time.Time
	Fn        Job
}

func (j *JobData) Exec() {
	j.StartTime = time.Now()
	j.Fn()
	j.EndTime = time.Now()
}

type JobsQueue chan *JobData
