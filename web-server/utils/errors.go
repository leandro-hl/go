package utils

import "net/http"

// SendError if error sends 500 over http
func SendError(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}

	return false
}

// ExecOrError Command is executed if no errors
func ExecOrError(w http.ResponseWriter, err error, fn Command) {
	if !SendError(w, err) {
		fn()
	}
}

func ExecOrExecError(err error, onErr Command, onSucc Command) {
	if err != nil {
		onErr()
	} else {
		onSucc()
	}
}

func ExecOnError(err error, onErr Command) {
	if err != nil {
		onErr()
	}
}
