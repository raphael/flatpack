package flatpack

import (
	"fmt"
	"reflect"
)

// BadType is an error that provides information about invalid types
// encountered while unmarshalling.
type BadType struct {
	Name   Key
	Kind   reflect.Kind
	reason string
}

func (e *BadType) Error() string {
	return fmt.Sprintf("flatpack: invalid type; %s (name=%s,kind=%s)", e.reason, e.Name, e.Kind)
}

// BadValue is an error that provides information about malformed values
// encountered while unmarshalling.
type BadValue struct {
	Name     Key
	Cause    error
	expected string
}

func (e *BadValue) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf(`flatpack: malformed value; (name=%s,cause="%s")`, e.Name, e.Cause.Error())
	}
	return fmt.Sprintf("flatpack: invalid value; expected %s (name=%s)", e.expected, e.Name)
}

// NoReflection is an error that indicates something went wrong when reflecting
// on an unmarshalling target. Generally, this is caused by trying to unmarshal
// into a struct that has unexported fields (i.e. whose names begin with a
// lower-case letter).
//
// To avoid this error, either remove the unexported fields from your struct
// or mark them with the flatpack:"ignore" field tag.
type NoReflection struct {
	Name Key
}

func (e *NoReflection) Error() string {
	return fmt.Sprintf("flatpack: reflection error; unexported field (name=%s)", e.Name)
}
