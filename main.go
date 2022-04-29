package main

import "fmt"

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
		return oo.Num.Add(d)

	case string:
		if oo.Str == nil {
			oo.Str = NewString()
		}
		return oo.Str.Add(d)

	case bool:
		if oo.Bool == nil {
			oo.Bool = NewBoolean()
		}
		return oo.Bool.Add(d)

	case map[string]any:
		if oo.Obj == nil {
			oo.Obj = NewObject()
		}
		return oo.Obj.Add(d)

	case []any:
		if oo.Arr == nil {
			oo.Arr = NewArray()
		}
		return oo.Arr.Add(d)

	case nil:
		if oo.Null == nil {
			oo.Null = NewNull()
		}
		return oo.Null.Add(d)

	}
	return fmt.Errorf("unknown type %T", a)
}

type Object struct {
	Count  int                     // Number of occurences of the object
	Fields map[string]*ObjectField // Mapping from object keys to object value schemas
}

func NewObject() *Object {
	return &Object{}
}

func (o *Object) Add(m map[string]any) error {
	// Add the object fields
	for k, v := range m {
		f := o.Fields[k]
		if f == nil {
			f = NewObjectField()
		}
		f.Add(v)
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
	err := arr.ItemType.Add(a)
	if err != nil {
		return err
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
	return nil
}

func main() {
	fmt.Println("Hello, world!")
}
