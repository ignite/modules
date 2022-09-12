// Package errors defines methods to handle specific error in cosmos blockchain like critical errors
package errors

import (
	"fmt"
)

const codespace = "CRITICAL"

var ErrCritical = Register(codespace, 2, "the state of the blockchain is inconsistent or an invariant is broken")

// Critical handles and/or returns an error in case a critical error has been encountered:
// - Inconsistent state
// - Broken invariant
func Critical(description string) error {
	return Wrap(ErrCritical, description)
}

// Criticalf extends a critical error with additional information.
//
// This function works like the Critical function with additional
// functionality of formatting the input as specified.
func Criticalf(format string, args ...interface{}) error {
	desc := fmt.Sprintf(format, args...)
	return Critical(desc)
}
