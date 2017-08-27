package jsonronin

import (
	"reflect"
)

const (
	Invalid int = iota
	Bool
	Number
	String
	Array
	Object
)

type Ronin struct {
	reflect.Value
}

func New(j interface{}) Ronin {
	return Ronin{reflect.ValueOf(j)}
}

func (v Ronin) Type() int {
	switch v.Kind() {
		case reflect.Float64:
			return Number

		case reflect.Bool:
			return Bool

		case reflect.String:
			return String

		case reflect.Slice, reflect.Array:
			return Array

		case reflect.Map:
			return Object

		default:
			return Invalid
	}	
}

func (v Ronin) ObjectItem(name string) Ronin {
	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			if(name == key.String()) {
				return Ronin{v.MapIndex(key).Elem()}
			}
		}	
	}
	return Ronin{}
}

func (v Ronin) ArrayItem(i int) Ronin {
	switch v.Kind() {
		case reflect.Slice, reflect.Array:
			return Ronin{v.Index(i).Elem()}
	}
	return Ronin{}
}

func (v Ronin) GetString() string {
	if v.Kind() == reflect.String {
		return v.String()
	}
	return ""
}

func (v Ronin) GetNumber() float64 {
	if v.Kind() == reflect.Float64 {
		return v.Float()
	}
	return 0
}

func (v Ronin) GetBool() bool {
	if v.Kind() == reflect.Bool {
		return v.Bool()
	}
	return false
}
