package main

import(
	"fmt"
	"net/http"
	"encoding/json"
)
var l = fmt.Println
const PORT = ":3000"

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/user", getuser)

	l("server listening 3000. http://127.0.0.1:3000/")
	err := http.ListenAndServe(PORT, nil)
	if err != nil { l(err) }
}




// module1
func hello(w http.ResponseWriter, req *http.Request) {
	data := "hello world"
	body(w,req,data)
}

// module 2
func getuser(w http.ResponseWriter, req *http.Request) {
	var user Users
	user.add("zs",0,18)
	user.add("ls",0,20)
	user.add("wmz",0,22)
	
	body(w,req,user.datas)
}




func body(w http.ResponseWriter, req *http.Request, body interface{})  {
	var(
		str string
		num int
	)
	str,sok := body.(string)
	num,nok := body.(int)
	if sok {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(str))
	}else if nok {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(string(num)))
	}else{
		w.Header().Set("Content-Type", "application/json")
		data,_ := json.Marshal(&body)
		w.Write(data)
	}
	
}


type Users struct{
	datas []map[string]interface{}
}

func (this *Users) add(name string,gender int,age int) {
	user := make(map[string]interface{})
	user["name"] = name
	user["gender"] = gender
	user["age"] = age
	this.datas = append(this.datas,user)
}


