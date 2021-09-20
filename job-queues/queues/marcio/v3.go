package marcio

import (
	"fmt"
	"log"
	"net/http"
)

type Job struct {
	Payload Payload
}

type JobQueue chan *Job

type Worker struct {
	ID       int
	JobQueue JobQueue
	quit     chan bool
}

func NewWorker(id int, jq JobQueue) *Worker {
	return &Worker{ID: id, JobQueue: jq, quit: make(chan bool)}
}

func (w *Worker) Start() {
	go func() {
		for {
			select {
			case job := <-w.JobQueue:
				fmt.Printf("Worker %v Working...\n", w.ID)
				if err := job.Payload.UploadToS3(); err != nil {
					log.Printf("Error uploading to S3: %s", err.Error())
				}
			case <-w.quit:
				// signal to stop processing
				// we return to break for loop and finish goroutine
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	//calling the channel in a goroutine so we don't block
	//current runtime
	go func() {
		w.quit <- true
	}()
}

var jobQueue JobQueue

func Init(workersCount *int, queueSize *int) {
	//create job queue
	jobQueue = make(JobQueue, *queueSize)

	//create and run workers
	for i := 1; i <= *workersCount; i++ {
		w := NewWorker(i, jobQueue)
		w.Start()
	}
}

func payloadHandlerV3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Payload received.")

	ct := ValidateAndInitialize2(w, r)

	if ct == nil {
		return
	}

	v3Processing(ct.Payloads)
}

func v3Processing(payloads []*Payload) {
	for _, payload := range payloads {
		job := &Job{*payload}

		//send to job queue
		jobQueue <- job
	}
}
