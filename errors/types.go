package errors

import "encoding/json"

type (
	Error struct {
		Code      string `json:"code"`
		Message   string `json:"message"`
		Details   string `json:"details,omitempty"`
		Field     string `json:"field,omitempty"`
		TraceID   string `json:"trace_id,omitempty"`
		Timestamp string `json:"timestamp,omitempty"`
	}
	FieldError struct {
		Message string `json:"message"`
		Field   string `json:"field,omitempty"`
	}
	ValidationError struct {
		Code      string       `json:"code"`
		Message   string       `json:"message"`
		Fields    []FieldError `json:"fields"`
		TraceID   string       `json:"trace_id,omitempty"`
		Timestamp string       `json:"timestamp,omitempty"`
	}
)

// Error implements error.
func (v ValidationError) Error() string {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return v.Message
	}
	return string(jsonData)
}

// Error implements error.
func (a Error) Error() string {
	jsonData, err := json.Marshal(a)
	if err != nil {
		return a.Message
	}
	return string(jsonData)
}

const (
	BAD_REQUEST       = "BAD_REQUEST"
	VALIDATION_FAILED = "VALIDATION_FAILED"
	ALREADY_EXISTS    = "ALREADY_EXISTS"
	NOT_FOUND         = "NOT_FOUND"
	UNAUTHORIZED      = "UNAUTHORIZED"
	FORBIDDEN         = "FORBIDDEN"
	INTERNAL_ERROR    = "INTERNAL_ERROR"
	DEPENDENCY_FAILED = "DEPENDENCY_FAILED"
	CONFLICT          = "CONFLICT"
	RATE_LIMITED      = "RATE_LIMITED"
	UNPROCESSABLE     = "UNPROCESSABLE"
	EXPIRED           = "EXPIRED"
	NOT_IMPLEMENTED   = "NOT_IMPLEMENTED"
)
