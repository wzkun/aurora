package errstring

// nsq and matterplatform

import (
	"fmt"
)

// ChildErr interface
type ChildErr interface {
	Parent() error
}

// ClientErr struct
type ClientErr struct {
	ParentErr     error
	Code          string `json:"code"`
	Desc          string `json:"desc"`           // Message to be display to the end user without debugging information
	DetailedError string `json:"detailed_error"` // Internal error string to help the developer
	Where         string `json:"-"`              // The function where it happened in the form of Struct.Func
	params        map[string]interface{}
}

func (e *ClientErr) Error() string {
	code := fmt.Sprintf("%s|%s|%s", e.Code, e.Where, e.DetailedError)
	return code
}

// Parent error
func (e *ClientErr) Parent() error {
	return e.ParentErr
}

// NewClientErr creates a ClientErr with the supplied human and machine readable strings
func NewClientErr(parent error, code, detail, where string, params map[string]interface{}) *ClientErr {
	ap := &ClientErr{}
	ap.ParentErr = parent
	ap.Code = code
	ap.Desc = code
	ap.DetailedError = detail
	ap.Where = where
	ap.params = params

	return ap
}
