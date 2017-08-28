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
	jr := jsonronin.New(v)

	#or
	jr :=jsonronin.Unmarshal(JsonStr)

#2 get string 

	name := jr.ObjectItem("name").String()
	
#3 get number

	age := jr.ObjectItem("age").Number()
	#or
	age := jr.ObjectItem("age").Float()
	
#4 get bool

	human := jr.ObjectItem("human").Bool()
	
#4 get object

	obj := jr.ObjectItem("object")
	obj := jr.ObjectItem("array")
	
#5 get array item

	item := obj.ArrayItem(1).String()
	item := obj.ArrayItem(1).Number()
	item := obj.ArrayItem(1).Bool()
	
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
		jr :=jsonronin.Unmarshal(JsonStr)
	
		name = jr.ObjectItem("name").String()
		fmt.Println(name)
		
		age := jr.ObjectItem("age").Number()
		fmt.Println(age)
	
		human  = jr.ObjectItem("human").Bool()
		fmt.Println(human)
		
		oobj := jr.ObjectItem("object")
		fmt.Println(oobj)
		aobj := jr.ObjectItem("array")
		fmt.Println(aobj)
		
		item  = aobj.ArrayItem(1).String()
		fmt.Println(item)
	}

#ouptut

	ronin
	24
	true
	<map[string]interface {} Value>
	<[]interface {} Value>
	b
