// Package logfilter provide support for filtering sensitive data defined in your structs from being written in your logs
package logfilter

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Call takes respectively as source a pointer to struct and as dest a pointer to string for filtering, returns error if any
func Call(raw interface{}, flt interface{}) error {
	if raw == nil || flt == nil {
		return errors.New("Only valid pointers accepted")
	}

	tofRaw := reflect.TypeOf(raw)
	tofFlt := reflect.TypeOf(flt)
	if tofRaw.Kind() != reflect.Ptr || tofFlt.Kind() != reflect.Ptr {
		return errors.New("Supported input and output for filtering are pointers")
	}

	rawKind := tofRaw.Elem().Kind()
	fltKind := tofFlt.Elem().Kind()
	if rawKind != reflect.Struct || fltKind != reflect.String {
		return errors.New("Supported input and output for filtering are respectivelly pointer to a struct and pointer to a string")
	}

	fltOut := reflect.Indirect(reflect.ValueOf(flt))
	trElem := tofRaw.Elem()

	pload := reflect.Indirect(reflect.ValueOf(raw))
	for i := 0; i < trElem.NumField(); i++ {
		f := trElem.Field(i)
		v, _ := f.Tag.Lookup("log")

		switch v {
		case "omit":
			continue
		case "filtered":
			fltOut.SetString(fmt.Sprintf("%s%s: %s; ", fltOut, trElem.Field(i).Name, "*************"))
		default:
			fltOut.SetString(fmt.Sprintf("%s%s: %s; ", fltOut, trElem.Field(i).Name, pload.Field(i)))
		}
	}
	fltOut.SetString(strings.TrimSpace(fltOut.String()))
	return nil
}
