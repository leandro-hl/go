package marcio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PayloadCollection struct {
	WindowsVersion *string   `json:"version"`
	Token          *string   `json:"token"`
	Payloads       []Payload `json:"payloads"`
}

type PayloadCollection2 struct {
	WindowsVersion *string    `json:"version"`
	Token          *string    `json:"token"`
	Payloads       []*Payload `json:"payloads"`
}

type Payload struct {
	Data *string `json:"data"`
}

func NewPayloadCollection() *PayloadCollection2 {
	return &PayloadCollection2{Payloads: []*Payload{}}
}

func (p *Payload) UploadToS3() error {
	//buffer := new(bytes.Buffer)
	//
	//err := json.NewEncoder(buffer).Encode(p)
	//
	//if err != nil {
	//	return err
	//}

	//SIMULATING TIME CONSUMING OPERATION
	fmt.Printf("Uploading payload: %v...\n", *p.Data)
	time.Sleep(2 * time.Second)
	fmt.Printf("Payload %v uploaded!\n", *p.Data)

	// LOOK FOR A NEWEST LIBRARY APPROACH TO UPLOAD TO AN S3 BUCKET
	// THIS EXAMPLE IS 6y OLD AND NOT WORKING ANYMORE

	//storagePath := fmt.Sprintf("%v/%v", p.storageFolder, time.Now().UnixNano())
	//
	////bucket here
	//bucket := s3.Bucket{}
	//
	//// Everything we post to the S3 bucket should be marked as private
	//accessControlList := s3.BucketCannedACLPrivate
	//contentType := "application/octet-stream"
	//return bucket.PutReader(storage_path, b, int64(b.Len()), contentType, acl, s3.Options{})

	return nil
}

// storageFolder ensures that there are no name collision in
// case we get same timestamp in the key name
func (p *Payload) storageFolder() string {
	return ""
}

func payloadHandler(w http.ResponseWriter, r *http.Request) {
	content := ValidateAndInitialize(w, r)

	if content != nil {
		return
	}

	v1Processing(content.Payloads)

	w.WriteHeader(http.StatusOK)
}

func ValidateAndInitialize(w http.ResponseWriter, r *http.Request) *PayloadCollection {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return nil
	}

	content := &PayloadCollection{}

	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		//tell we expect json
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
	}

	return content
}

func ValidateAndInitialize2(w http.ResponseWriter, r *http.Request) *PayloadCollection2 {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return nil
	}

	content := NewPayloadCollection()

	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		//tell we expect json
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
	}

	return content
}

// go through each payload and queue items individually to be
// posted to S3.
// this approach spawn a new go routine per request
// that means we're launching them in an uncontrolled way.
// this is bad. In addition, we're finishing the request
// without knowing if the routine has finished or not
func v1Processing(payloads []Payload) {
	for _, payload := range payloads {
		go payload.UploadToS3()
	}
}
