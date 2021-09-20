package shared

import (
	"bufio"
	"log"
	"os"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

//Interface embedding.
//ReadWriter gets the methods defined on Reader and Writer directly
type IReadWriter interface {
	Reader
	Writer
}

//Struct embedding.
//ReadWriter holds pointers to both bufio reader and writer.
//their methods can be invoked directly from ReadWriter.
//To do this, do not name the fields. Embedding are unamed fields.
type ReadWriter struct {
	*bufio.Reader
	*bufio.Writer
	//No embedding approach
	reader *bufio.Reader
}

//No embedding approach
func (rw *ReadWriter) Read(p []byte) (n int, err error) {
	return rw.reader.Read(p)
}

func ReadWriterCheck() {
	rw := &ReadWriter{}

	b := []byte("Read Write check")

	//Using Reader methods
	rw.Read(b)

	//Using Writer methods
	rw.Write(b)
}

//Mixing embedding with a common named field
type Job struct {
	Command string
	*log.Logger
}

func TestJob() {
	j := &Job{"a command to exec", log.New(os.Stderr, "Job: ", log.Ldate)}

	j.Println("Using methods from log.Logger")

	//this Printf is from definition below, not from the Logger
	//removing below method declaration would allow to use the Logger Printf directly.
	//Kind of an "Override"
	j.Printf("%d")
}

//We can also access the embedded type by typing its name
func (j *Job) Printf(format string) {
	j.Logger.Printf(format, j.Command)
}
