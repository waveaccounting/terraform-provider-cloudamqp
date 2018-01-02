package cloudamqp

import (
	"fmt"
	"log"
)

// APIError represents a CloudAMQP API Error response
type APIError struct {
	Status int
}

func (e APIError) Error() string {
	if e.Status > 299 {
		fmt.Sprint("cloudamqp: non-200 status code response")
	}

	return fmt.Sprintf("cloudamqp: %v", e.Status)
}

// Empty returns true if empty.
func (e APIError) Empty() bool {
	return 200 <= e.Status && e.Status <= 299
}

func relevantError(httpError error, apiError APIError) error {
	if httpError != nil {
		log.Printf("%v", "Error = httpError")
		return httpError
	}
	log.Printf("apiError = %+v", apiError)

	if !apiError.Empty() {
		log.Printf("%v", "Error = apiError")
		return apiError
	}
	return nil
}
