package lakesite_test

import (
	"fmt"
	q "hl/job-queues/queues/lakesite"
	"testing"
)

func BenchmarkDispatcher_Startc(b *testing.B) {
	var jobs []q.Job

	for i := 0; i < 10; i++ {
		jobs = append(jobs, func() error {
			fmt.Println("Inside: Job 1 starting...")
			fmt.Println("Inside: Job 1 Finished.")
			return nil
		})
	}

	d := q.NewDispatcher(10)

	go func() {
		for _, j := range jobs {
			d.AddJob(j)
		}
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Start(2)
	}
}
