
package jsonronin

import (
	"testing"
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

	o := jr.ObjectItem("object")
	a := jr.ObjectItem("array")
	s := jr.ObjectItem("name")
	n := jr.ObjectItem("age")
	b := jr.ObjectItem("good")

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
	jr := Unmarshal(testJSON)

	o := jr.ObjectItem("object")
	if o.Type() != Object {
		t.Error("ObjectItem() != object")
	}
}
func TestArrayItem(t *testing.T) {
	jr := Unmarshal(testJSON)

	a := jr.ObjectItem("array")
	i := a.ArrayItem(1)
	if i.Number() != 2 {
		t.Error("ObjectItem(1) != 2")
	}
}
func TestString(t *testing.T) {
	jr := Unmarshal(testJSON)

	name := jr.ObjectItem("name").String()
	if name != "ronin" {
		t.Error("String(name) != ronin")
	}
}
func TestNumber(t *testing.T) {
	jr := Unmarshal(testJSON)

	age := jr.ObjectItem("age").Number()
	if age != 100 {
		t.Error("Number(age) != 100")
	}
}
func TestBool(t *testing.T) {
	jr := Unmarshal(testJSON)

	bt := jr.ObjectItem("good").Bool()
	if bt != true {
		t.Error("Bool(good) != true")
	}
	bf := jr.ObjectItem("bad").Bool()
	if bf != false {
		t.Error("Bool(bad) != true")
	}
}