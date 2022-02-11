package main

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"time"

	httpErr "github.com/Yuichi-Kadota/go_error_handle/internal/error"
)

func main() {
	// generate rand(0->4)
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(5)
	err := Process(i)
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

func Process(i int) error {
	switch i {
	case 0:
		return InternalErrorProcess()
	case 1:
		return UnAuthorizedErrorProcess()
	case 2:
		return ForbiddenErrorProcess()
	case 3:
		return UnControlledErrorProcess()
	default:
		return nil
	}
}

func InternalErrorProcess() error {
	return &httpErr.InternalServerErr{
		Origin: errors.New("internal server error"),
	}
}

func UnAuthorizedErrorProcess() error {
	return &httpErr.UnAuthorized{
		Origin: errors.New("unAuthorized error"),
	}
}

func ForbiddenErrorProcess() error {
	return &httpErr.Forbidden{
		Origin: errors.New("forbiden error"),
	}
}

func UnControlledErrorProcess() error {
	return errors.New("unControlled Err")
}
