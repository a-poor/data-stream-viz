package dsviz

import (
	"fmt"
)

// OneOf represents the different possible types for a field.
type OneOf struct {
	Obj  *Object  `json:"obj,omitempty"`   // Schema of object values
	Arr  *Array   `json:"arr,omitempty"`   // Schema of array values
	Num  *Number  `json:"num,omitempty"`   // Schema of number values
	Str  *String  `json:"str,omitempty"`   // Schema of string values
	Bool *Boolean `json:"bool,omitempty"`  // Schema of boolean values
	Null *Null    `json:"null_,omitempty"` // Schema of null values
}

// Add updates the schema with the given data.
func (oo *OneOf) Add(a any) error {
	switch d := a.(type) {
	case float64:
		if oo.Num == nil {
			oo.Num = NewNumber()
		}
		err := oo.Num.Add(d)
		return addPathToError(err, ".(number)")

	case string:
		if oo.Str == nil {
			oo.Str = NewString()
		}
		err := oo.Str.Add(d)
		return addPathToError(err, ".(string)")

	case bool:
		if oo.Bool == nil {
			oo.Bool = NewBoolean()
		}
		err := oo.Bool.Add(d)
		return addPathToError(err, ".(boolean)")

	case map[string]any:
		if oo.Obj == nil {
			oo.Obj = NewObject()
		}
		err := oo.Obj.Add(d)
		return addPathToError(err, ".(object)")

	case []any:
		if oo.Arr == nil {
			oo.Arr = NewArray()
		}
		err := oo.Arr.Add(d)
		return addPathToError(err, ".(array)")

	case nil:
		if oo.Null == nil {
			oo.Null = NewNull()
		}
		err := oo.Null.Add(d)
		return addPathToError(err, ".(null)")

	}
	return fmt.Errorf("unknown type %T", a)
}

// Object represents a schema for an object.
type Object struct {
	Count  int                     `json:"count"`  // Number of occurences of the object
	Fields map[string]*ObjectField `json:"fields"` // Mapping from object keys to object value schemas
}

// NewObject creates a new Object schema.
func NewObject() *Object {
	return &Object{
		Fields: make(map[string]*ObjectField),
	}
}

// Add updates the object schema with new object data
func (o *Object) Add(m map[string]any) error {
	// Add the object fields
	for k, v := range m {
		f := o.Fields[k]
		if f == nil {
			f = NewObjectField()
		}
		err := f.Add(v)
		if err != nil {
			return addPathToError(err, fmt.Sprintf(".%s", k))
		}
		o.Fields[k] = f
	}

	// Increment the counter
	o.Count += 1
	return nil
}

// ObjectFirld represents a field in a JSON object.
type ObjectField struct {
	Count int   `json:"count"` // Number of occurences of the object field
	Type  OneOf `json:"type"`  // Schema of the object field
}

// NewObjectFirld creates a new ObjectField schema.
func NewObjectField() *ObjectField {
	return &ObjectField{}
}

// Add updates the object firld schema with new object field data.
func (of *ObjectField) Add(a any) error {
	err := of.Type.Add(a)
	if err != nil {
		return err
	}
	of.Count += 1
	return nil
}

// Array represents a JSON array's schema.
type Array struct {
	Count    int   // Number of occurences of the array object
	ItemType OneOf // Schema of the array items
}

// NewArray creates a new Array schema.
func NewArray() *Array {
	return &Array{}
}

// Add updates the array schema with new array data.
func (arr *Array) Add(a []any) error {
	for _, d := range a {
		err := arr.ItemType.Add(d)
		if err != nil {
			return addPathToError(err, "[]")
		}
	}
	arr.Count += 1
	return nil
}

// Number represents a JSON number's schema.
//
// Note: This assumes that the JSON object's number was unmarshalled as a float64
// which is the normal behavior for the "json" package.
type Number struct {
	Count int `json:"count"` // Number of occurences of the number value
}

// NewNumber creates a new Number schema.
func NewNumber() *Number {
	return &Number{}
}

// Add updates the number schema with new number data.
//
// Note: This assumes that the JSON object's number was unmarshalled as a float64
// which is the normal behavior for the "json" package.
func (n *Number) Add(float64) error {
	n.Count += 1
	return nil
}

// String represents a JSON string's schema.
type String struct {
	Count int `json:"count"` // Number of occurences of the string value
}

// NewString creates a new String schema.
func NewString() *String {
	return &String{}
}

// Add updates the string schema with new string data.
func (s *String) Add(string) error {
	s.Count += 1
	return nil
}

// Boolean represents a JSON boolean's schema.
type Boolean struct {
	Count int `json:"count"` // Number of occurences of the bool value
}

// NewBoolean creates a new Boolean schema.
func NewBoolean() *Boolean {
	return &Boolean{}
}

// Add updates the boolean schema with new bool data.
func (b *Boolean) Add(bool) error {
	b.Count += 1
	return nil
}

// Null represents a JSON null's schema.
type Null struct {
	Count int `json:"count"` // Number of occurences of the null value
}

// NewNull creates a new Null schema.
func NewNull() *Null {
	return &Null{}
}

// Add updates the null schema with new null data.
func (n *Null) Add(a any) error {
	if a != nil {
		return fmt.Errorf("expected type nil, got %T", a)
	}
	n.Count += 1

	return nil
}
