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
			100,
			201,
			"TOP"
		]
	}`
#1 init 

	var v interface{}
	json.Unmarshal([]byte(JsonStr), &v)
	jr := jsonronin.New(v)

	#or
	jr :=jsonronin.Unmarshal(JsonStr)

#2 get string 

	name := jr.Get("name").String()
	
#3 get number

	age := jr.Get("age").Number()
	#or
	age := jr.Get("age").Float()
	
#4 get bool

	human := jr.Get("human").Bool()
	
#4 get object item

	obj := jr.Get("object")
	itemx := jr.Get("object").Get("x").Number()
	itemx := jr.Get("object").Get("y").String()
	
#5 get array item

	list  := jr.Get("array")

	item0 := jr.Get("array").Get("0").String()
	item1 := jr.Get("array").Get("1").Number()
	item2 := jr.Get("array").Get("2").String()
	
#end

#example
