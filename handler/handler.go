package handler

import (
	"fmt"
	"net/http"

	"github.com/Aritiaya50217/CodeBangkok/errs"
)

func handlerError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case errs.AppError:
		w.WriteHeader(e.Code)
		fmt.Println(w, e)
	case error:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(fmt.Fprintln(w, e))
	}
}
