package errors

func NewError(code, message string, options ...func(*Error)) Error {
	err := Error{
		Code:    code,
		Message: message,
	}

	for _, opt := range options {
		opt(&err)
	}

	return err
}

func NewFieldError(field, message string) FieldError {
	return FieldError{
		Field:   field,
		Message: message,
	}
}

func NewValidationError(fields []FieldError, options ...func(*ValidationError)) ValidationError {
	err := ValidationError{
		Code:    VALIDATION_FAILED,
		Message: "Validation failed on one or more fields.",
		Fields:  fields,
	}

	for _, opt := range options {
		opt(&err)
	}

	return err
}

func WithTraceID(traceID string) func(interface{}) {
	return func(e interface{}) {
		switch v := e.(type) {
		case *Error:
			v.TraceID = traceID
		case *ValidationError:
			v.TraceID = traceID
		}
	}
}

func WithDetails(details string) func(*Error) {
	return func(e *Error) {
		e.Details = details
	}
}

func WithField(field string) func(*Error) {
	return func(e *Error) {
		e.Field = field
	}
}

func WithTimestampApi(ts string) func(*Error) {
	return func(e *Error) {
		e.Timestamp = ts
	}
}

func WithTimestampValidation(ts string) func(*ValidationError) {
	return func(e *ValidationError) {
		e.Timestamp = ts
	}
}
