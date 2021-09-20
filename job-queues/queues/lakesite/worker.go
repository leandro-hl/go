package lakesite

import (
	"fmt"
	"log"
)

//a. Register itself with the dispatcher's channel of workers.
//b. Process jobs from the job queue channel.
//c. Respond to orders from the dispatcher to halt or get job status.
//d. Set properties on a job (start time, end time, result).
type Worker struct {
	ID       int
	Finished chan bool
	jobs     JobsQueue
	l        *log.Logger
}

func NewWorker(id int, jobs JobsQueue, finished chan bool) *Worker {
	return &Worker{ID: id, jobs: jobs, Finished: finished}
}

func (w *Worker) Start() {
	go func() {
		for {
			select {
			case job := <-w.jobs:
				w.Starting(job.ID)

				job.Exec()

				w.Finishing(job.ID)
			}
		}
	}()
}

func (w *Worker) Starting(jobId int) {
	fmt.Printf("Worker: %d starting job: %d.\n", w.ID, jobId)
}

func (w *Worker) Finishing(jobId int) {
	fmt.Printf("Worker: %d job: %d Finished.\n", w.ID, jobId)

	w.Finished <- true
}
