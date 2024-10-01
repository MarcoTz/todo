package types 

import ( 
  "fmt"
)

type DecodingErr struct {
  field_name string 
  message string
}
func (e *DecodingErr) Error() string { return fmt.Sprintf("Error deconding response, field: %s, message :%s",e.field_name,e.message) }

type BadRequestErr struct { 
  contents string
}
func (e *BadRequestErr) Error() string { return fmt.Sprintf("Got error response:\n%s",e.contents) }
