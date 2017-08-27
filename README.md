# jsonronin

var JsonStr = `{<br>
	"name":"ronin",<br>
	"age":24,<br>
	"human":true,<br>
	"object":{<br>
		"x":1,<br>
		"y":"bbb"<br>
	},<br>
	"array":[<br>
		"a",<br>
		"b",<br>
		"c"<br>
	]<br>
}`<br>

#1 init

	var v interface{}
	json.Unmarshal([]byte(JsonStr), &v)

	#json ronin
	jr := jsonronin.New(v)


#2 get string <br>
name := jr.ObjectItem("name").GetString()<br>
#or<br>
name := jr.ObjectItem("name").String()<br>
<br>
#3 get number<br>
age := jr.ObjectItem("age").GetNumber()<br>
#or<br>
age := jr.ObjectItem("age").Float()<br>
<br>
#4 get bool<br>
human := jr.ObjectItem("human").GetBool()<br>
#or<br>
human := jr.ObjectItem("human").Bool()<br>
<br>
#4 get object<br>
obj := jr.ObjectItem("object")<br>
obj := jr.ObjectItem("array")<br>
<br>
#5 get array item<br>
item := obj.ArrayItem(1).GetString()<br>
item := obj.ArrayItem(1).GetNumber()<br>
item := obj.ArrayItem(1).GetBool()<br>
<br>
#array[object,object....]<br>
subobj := obj.ArrayItem(1)<br>
<br>
<br>
#end<br>

#example<br>
<br>
package main<br>
<br>
import (<br>
	"fmt"<br>
	"encoding/json"<br>
	"github.com/gaseity/jsonronin"<br>
)<br>
<br>
var JsonStr = `{<br>
	"name":"ronin",<br>
	"age":24,<br>
	"human":true,<br>
	"object":{<br>
		"x":1,<br>
		"y":"bbb"<br>
	},<br>
	"array":[<br>
		"a",<br>
		"b",<br>
		"c"<br>
	]<br>
}`<br>
<br>
func main() {<br>
	var v interface{}<br>
	json.Unmarshal([]byte(JsonStr), &v)<br>
	jr := jsonronin.New(v)<br>

<br>
name := jr.ObjectItem("name").GetString()<br>
fmt.Println(name)<br>
name = jr.ObjectItem("name").String()<br>
fmt.Println(name)<br>
<br>
age := jr.ObjectItem("age").GetNumber()<br>
fmt.Println(age)<br>
age = jr.ObjectItem("age").Float()<br>
fmt.Println(age)<br>
<br>
human := jr.ObjectItem("human").GetBool()<br>
fmt.Println(human)<br>
human  = jr.ObjectItem("human").Bool()<br>
fmt.Println(human)<br>
<br>
oobj := jr.ObjectItem("object")<br>
fmt.Println(oobj)<br>
aobj := jr.ObjectItem("array")<br>
fmt.Println(aobj)<br>
<br>
item := aobj.ArrayItem(0).GetString()<br>
fmt.Println(item)<br>
item  = aobj.ArrayItem(1).String()<br>
fmt.Println(item)<br>
<br>
}<br>

#ouptut<br>
ronin<br>
ronin<br>
24<br>
24<br>
true<br>
true<br>
<map[string]interface {} Value><br>
<[]interface {} Value><br>
a<br>
b<br>
