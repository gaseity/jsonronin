
package jsonronin

import (
	"encoding/json"
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

func ready(j string) interface{} {
	var v interface{}
	json.Unmarshal([]byte(j), &v)
	return v;
}

func TestType(t *testing.T) {
	jr := New(ready(testJSON))

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
	jr := New(ready(testJSON))

	o := jr.ObjectItem("object")
	if o.Type() != Object {
		t.Error("ObjectItem() != object")
	}
}
func TestArrayItem(t *testing.T) {
	jr := New(ready(testJSON))

	a := jr.ObjectItem("array")
	i := a.ArrayItem(1)
	if i.GetNumber() != 2 {
		t.Error("ObjectItem(1) != 2")
	}
}
func TestGetString(t *testing.T) {
	jr := New(ready(testJSON))

	name := jr.ObjectItem("name").GetString()
	if name != "ronin" {
		t.Error("String(name) != ronin")
	}
}
func TestGetNumber(t *testing.T) {
	jr := New(ready(testJSON))

	age := jr.ObjectItem("age").GetNumber()
	if age != 100 {
		t.Error("Number(age) != 100")
	}
}
func TestGetBool(t *testing.T) {
	jr := New(ready(testJSON))

	bt := jr.ObjectItem("good").GetBool()
	if bt != true {
		t.Error("Bool(good) != true")
	}
	bf := jr.ObjectItem("bad").GetBool()
	if bf != false {
		t.Error("Bool(bad) != true")
	}
}