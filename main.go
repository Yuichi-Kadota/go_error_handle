package main

import (
	"errors"
	"log"
	"net/http"

	httpErr "github.com/Yuichi-Kadota/go_error_handle/internal/error"
)

func main() {
	err := Process()
	if err != nil {
		// Get Status Code
		result, ok := err.(httpErr.HttpErr)
		if ok {
			log.Println(result.Code())
			return
		}
		log.Printf("unexpected error %v", err)
		return
	}
	log.Println(http.StatusOK)
}

func Process() error {
	err := InternalErrorProcess()
	//err := UnControlledErrorProcess()
	if err != nil {
		return err
	}
	return nil
}

func InternalErrorProcess() error {
	return httpErr.InternalServerErr{
		Origin: errors.New("internal server error"),
	}
}

func UnAuthorizedErrorProcess() error {
	return httpErr.UnAuthorized{
		Origin: errors.New("unAuthorized error"),
	}
}

func ForbiddenErrorProcess() error {
	return httpErr.Forbidden{
		Origin: errors.New("forbiden error"),
	}
}

func UnControlledErrorProcess() error {
	return errors.New("unControlled Err")
}
