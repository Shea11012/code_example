// Code generated by entc, DO NOT EDIT.

package ent

import (
	"shor_url/ent/schema"
	"shor_url/ent/tinyurl"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	tinyurlFields := schema.TinyURL{}.Fields()
	_ = tinyurlFields
	// tinyurlDescID is the schema descriptor for id field.
	tinyurlDescID := tinyurlFields[0].Descriptor()
	// tinyurl.IDValidator is a validator for the "id" field. It is called by the builders before save.
	tinyurl.IDValidator = tinyurlDescID.Validators[0].(func(uint64) error)
}
