package jsonronin

import (
	"strconv"
	"encoding/json"
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

func Unmarshal(s string) Ronin {
	var v interface{}
	json.Unmarshal([]byte(s), &v)
	return New(v)
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

func (v Ronin) Pick(i int) Ronin {
	switch v.Kind() {

	case reflect.Slice, reflect.Array:
		return Ronin{v.Index(i).Elem()}
	}
	return Ronin{}
}

func (v Ronin) Get(k string) Ronin {
	switch v.Kind() {

	case reflect.Slice, reflect.Array:
		i,_ := strconv.Atoi(k)
		return Ronin{v.Index(i).Elem()}

	case reflect.Map:
		for _, key := range v.MapKeys() {
			if(k == key.String()) {
				return Ronin{v.MapIndex(key).Elem()}
			}
		}	
	}
	return Ronin{}
}

func (v Ronin) Interface() interface{} {
	return v.Interface()	
}

func (v Ronin) Number() float64 {
	return v.Float()
}
