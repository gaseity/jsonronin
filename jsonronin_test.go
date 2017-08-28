
package jsonronin

import (
	"testing"
	"fmt"
)

var testJSON = `{
	"age":100, 
	"name":"ronin",
	"object":{"what is a this?":"an apple"},
	"good":true,
	"bad":false,
	"array":[1,2,3],
	"items":["a","b","c"]
}`

func TestType(t *testing.T) {
	jr := Unmarshal(testJSON)

	o := jr.Get("object")
	a := jr.Get("array")
	s := jr.Get("name")
	n := jr.Get("age")
	b := jr.Get("good")

	if o.Type() != Object {
		t.Error("object.Type() != object")
	}
	if a.Type() != Array {
		t.Error("array.Type() != array")
	}
	if s.Type() != String {
		t.Error("string.Type() != string")
	}
	if n.Type() != Number {
		t.Error("number.Type() != number")
	}
	if b.Type() != Bool {
		t.Error("bool.Type() != bool")
	}
}

func TestObjectItem(t *testing.T) {
	jr := Unmarshal(`{"object":{"what is a this?":"an apple"}}`)

	o := jr.Get("object")
	if o.Type() != Object {
		t.Error("Object() != object")
	}
	ow := jr.Get("object").Get("what is a this?").String()
	if ow != "an apple" {
		t.Error("Object(item) != \"an apple\"")
	}
}
func TestArray(t *testing.T) {
	jr := Unmarshal(`{"array":[0,1,"2"]}`)

	a := jr.Get("array").Get("0").Number()
	if a != 0 {
		t.Error("Array(0) != 1")
	}
	b := jr.Get("array").Get("1").Number()
	if b != 1 {
		t.Error("Array(1) != 2")
	}
	c := jr.Get("array").Get("2").String()
	if c != "2" {
		t.Error("Array(2) != \"2\"")
	}
}
func TestString(t *testing.T) {
	jr := Unmarshal(`{"name":"ronin"}`)

	name := jr.Get("name").String()
	if name != "ronin" {
		t.Error("String(name) != ronin")
	}

}
func TestNumber(t *testing.T) {
	jr := Unmarshal(`{"age":100}`)

	age := jr.Get("age").Number()
	if age != 100 {
		t.Error("Number(age) != 100")
	}
}
func TestBool(t *testing.T) {
	jr := Unmarshal(`{"good":true, "bad":false}`)

	bt := jr.Get("good").Bool()
	if bt != true {
		t.Error("Bool(good) != true")
	}
	bf := jr.Get("bad").Bool()
	if bf != false {
		t.Error("Bool(bad) != true")
	}
}

