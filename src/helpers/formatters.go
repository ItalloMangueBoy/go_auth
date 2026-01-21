package helpers

import (
	"strings"
)

// FormatString formats a string to a more readable format
func FormatString(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}

// FormatOptionalString formats a optional string to a more readable format
func FormatOptionalString(str *string) *string {
	if str == nil {
		return nil
	}

	formmated := FormatString(*str)
	return &formmated
}
