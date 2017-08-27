# jsonronin

var JsonStr = `{
	"name":"ronin",
	"age":24,
	"human":true,
	"object":{
		"x":1,
		"y":"bbb"
	},
	"array":[
		"a",
		"b",
		"c"
	]
}`

#1 init

	var v interface{}
	json.Unmarshal([]byte(JsonStr), &v)

	#json ronin
	jr := jsonronin.New(v)


#2 get string 
name := jr.ObjectItem("name").GetString()
#or
name := jr.ObjectItem("name").String()

#3 get number
age := jr.ObjectItem("age").GetNumber()
#or
age := jr.ObjectItem("age").Float()

#4 get bool
human := jr.ObjectItem("human").GetBool()
#or
human := jr.ObjectItem("human").Bool()

#4 get object
obj := jr.ObjectItem("object")
obj := jr.ObjectItem("array")

#5 get array item
item := obj.ArrayItem(1).GetString()
item := obj.ArrayItem(1).GetNumber()
item := obj.ArrayItem(1).GetBool()

#array[object,object....]
subobj := obj.ArrayItem(1)


#end

#example

package main

import (
	"fmt"
	"encoding/json"
	"github.com/gaseity/jsonronin"
)

var JsonStr = `{
	"name":"ronin",
	"age":24,
	"human":true,
	"object":{
		"x":1,
		"y":"bbb"
	},
	"array":[
		"a",
		"b",
		"c"
	]
}`

func main() {
	var v interface{}
	json.Unmarshal([]byte(JsonStr), &v)
	jr := jsonronin.New(v)


name := jr.ObjectItem("name").GetString()
fmt.Println(name)
name = jr.ObjectItem("name").String()
fmt.Println(name)

age := jr.ObjectItem("age").GetNumber()
fmt.Println(age)
age = jr.ObjectItem("age").Float()
fmt.Println(age)

human := jr.ObjectItem("human").GetBool()
fmt.Println(human)
human  = jr.ObjectItem("human").Bool()
fmt.Println(human)

oobj := jr.ObjectItem("object")
fmt.Println(oobj)
aobj := jr.ObjectItem("array")
fmt.Println(aobj)

item := aobj.ArrayItem(0).GetString()
fmt.Println(item)
item  = aobj.ArrayItem(1).String()
fmt.Println(item)

}

#ouptut
ronin
ronin
24
24
true
true
<map[string]interface {} Value>
<[]interface {} Value>
a
b
