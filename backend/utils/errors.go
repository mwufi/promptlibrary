package utils

import (
	"fmt"
	"runtime"
	"strings"
)

// StackTrace represents a stack trace
type StackTrace struct {
	File     string
	Line     int
	Function string
}

// AppError wraps an error with stack trace information
type AppError struct {
	Err        error
	StackTrace []StackTrace
	Message    string
}

func (e *AppError) Error() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Error: %v\n", e.Message))
	if e.Err != nil {
		sb.WriteString(fmt.Sprintf("Cause: %v\n", e.Err))
	}
	sb.WriteString("Stack Trace:\n")
	for _, frame := range e.StackTrace {
		sb.WriteString(fmt.Sprintf("\t%s:%d %s\n", frame.File, frame.Line, frame.Function))
	}
	return sb.String()
}

// WrapError creates a new AppError with stack trace
func WrapError(err error, message string) *AppError {
	if err == nil {
		return nil
	}

	// If it's already an AppError, just update the message
	if appErr, ok := err.(*AppError); ok {
		appErr.Message = message
		return appErr
	}

	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(2, pcs[:])
	
	var trace []StackTrace
	frames := runtime.CallersFrames(pcs[:n])
	for {
		frame, more := frames.Next()
		
		// Skip runtime and standard library frames
		if !strings.Contains(frame.File, "runtime/") && !strings.Contains(frame.File, "/usr/local/go/") {
			trace = append(trace, StackTrace{
				File:     frame.File,
				Line:     frame.Line,
				Function: frame.Function,
			})
		}
		
		if !more {
			break
		}
	}

	return &AppError{
		Err:        err,
		StackTrace: trace,
		Message:    message,
	}
}

// NewError creates a new AppError without a cause
func NewError(message string) *AppError {
	return WrapError(fmt.Errorf(message), message)
}
