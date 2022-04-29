package dsviz

import (
	"fmt"
)

// OneOf represents different
type OneOf struct {
	Obj  *Object  // Schema of object values
	Arr  *Array   // Schema of array values
	Num  *Number  // Schema of number values
	Str  *String  // Schema of string values
	Bool *Boolean // Schema of boolean values
	Null *Null    // Schema of null values
}

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

type Object struct {
	Count  int                     // Number of occurences of the object
	Fields map[string]*ObjectField // Mapping from object keys to object value schemas
}

func NewObject() *Object {
	return &Object{
		Fields: make(map[string]*ObjectField),
	}
}

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

type ObjectField struct {
	Count int   // Number of occurences of the object field
	Type  OneOf // Schema of the object field
}

func NewObjectField() *ObjectField {
	return &ObjectField{}
}

func (of *ObjectField) Add(a any) error {
	err := of.Type.Add(a)
	if err != nil {
		return err
	}
	of.Count += 1
	return nil
}

type Array struct {
	Count    int   // Number of occurences of the array object
	ItemType OneOf // Schema of the array items
}

func NewArray() *Array {
	return &Array{}
}

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

type Number struct {
	Count int // Number of occurences of the number value
}

func NewNumber() *Number {
	return &Number{}
}

func (n *Number) Add(float64) error {
	n.Count += 1
	return nil
}

type String struct {
	Count int // Number of occurences of the string value
}

func NewString() *String {
	return &String{}
}

func (s *String) Add(string) error {
	s.Count += 1
	return nil
}

type Boolean struct {
	Count int // Number of occurences of the bool value
}

func NewBoolean() *Boolean {
	return &Boolean{}
}

func (b *Boolean) Add(bool) error {
	b.Count += 1
	return nil
}

type Null struct {
	Count int // Number of occurences of the null value
}

func NewNull() *Null {
	return &Null{}
}

func (n *Null) Add(a any) error {
	if a != nil {
		return fmt.Errorf("expected type nil, got %T", a)
	}
	n.Count += 1

	return fmt.Errorf("a fake error")
	// return nil
}
