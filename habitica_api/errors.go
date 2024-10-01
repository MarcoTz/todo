package habitica_api

import (
	"fmt"
)

type MissingVarErr struct { variable string } 
func (e *MissingVarErr) Error() string { return fmt.Sprintf("Missing environment variable %s",e.variable) } 

type CreateRequestErr struct{ message string }

func (e *CreateRequestErr) Error() string {
	return fmt.Sprintf("Error creating request: %s", e.message)
}

type PerformRequestErr struct{ message string }

func (e *PerformRequestErr) Error() string {
	return fmt.Sprintf("Error performing request: %s", e.message)
}

type ReaderErr struct {
	reading string
	message string
}

func (e *ReaderErr) Error() string { return fmt.Sprintf("Error reading %s: %s", e.reading, e.message) }

type ResponseErr struct { 
  code int 
  message string 
}

func (e *ResponseErr) Error() string { return fmt.Sprintf("Error performing request, status code %d, message: %s",e.code,e.message) }  
