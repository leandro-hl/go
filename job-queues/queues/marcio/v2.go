package marcio

import "net/http"

// Creating a buffered channel we can control how many
// items are there queued, that way reusing workers which
// are listening on different go routines
type Queue chan Payload

var queue Queue

const (
	MAX_QUEUED_ITEMS int = 10
)

func InitV2() {
	queue = make(Queue, MAX_QUEUED_ITEMS)
	processor()
}

func payloadHandlerV2(w http.ResponseWriter, r *http.Request) {
	content := ValidateAndInitialize(w, r)

	if content != nil {
		return
	}

	v2Processing(content.Payloads)

	w.WriteHeader(http.StatusOK)
}

func v2Processing(payloads []Payload) {
	for _, payload := range payloads {
		queue <- payload
	}
}

// "traded flawed concurrency with a buffered queue that was simply postponing the problem"
// "Our synchronous processor was only uploading one payload at a time to S3,
// and since the rate of incoming requests were much larger than the ability of
// the single processor to upload to S3, our buffered channel was quickly
// reaching its limit and blocking the request handler ability to queue
// more items."
func processor() {
	for {
		select {
		case payload := <-queue:
			//still spawning a go routine peer payload
			//means N peer request!!
			go payload.UploadToS3()
		}
	}
}
