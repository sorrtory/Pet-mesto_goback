// Chat GPT like parses the string...

package common

import (
	"fmt"
	"strings"
)

// Error structure for individual field errors
type FieldError struct {
	Name       string     `json:"name"`
	Message    string     `json:"message"`
	Properties Properties `json:"properties"`
	Kind       string     `json:"kind"`
	Path       string     `json:"path"`
}

// Properties structure for additional error details
type Properties struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Path    string `json:"path"`
}

// Final structured error response
type ValidationErrorResponse struct {
	Errors   map[string]FieldError `json:"errors"`
	Message  string                `json:"message"`
	Name     string                `json:"name"`
	Message2 string                `json:"_message"`
}

// Function to format validation errors
func FormatValidationError(err error) ValidationErrorResponse {
	validationErrors := make(map[string]FieldError)

	// Parse error string (this example assumes err.Error is in a specific format)
	errors := strings.Split(err.Error(), "\n")
	for _, e := range errors {
		// Example error: "Key: 'UserMe.Name' Error:Field validation for 'Name' failed on the 'required' tag"
		if e == "" {
			continue
		}

		// Extract the field and the error message
		parts := strings.Split(e, "Error:Field validation for ")
		if len(parts) == 2 {
			fieldName := strings.TrimSpace(parts[1])
			field := strings.Trim(strings.Split(fieldName, "'")[1], "'")
			message := fmt.Sprintf("%s cannot be empty", field)

			validationErrors[field] = FieldError{
				Name:    "ValidatorError",
				Message: message,
				Properties: Properties{
					Message: message,
					Type:    "required",
					Path:    field,
				},
				Kind: "required",
				Path: field,
			}
		}
	}

	// Construct the overall error response
	return ValidationErrorResponse{
		Errors:   validationErrors,
		Message:  "User validation failed: name: name cannot be empty, about: about cannot be empty",
		Name:     "ValidationError",
		Message2: "User validation failed",
	}
}
