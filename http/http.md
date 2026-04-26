package main

// for handler

import (
	"fmt"
	"log"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err:= r.ParseForm()
	if err != nil{
		log.Fatalln(err)
	}
	
	fmt.Fprintln(w, "Any code you want to perform in this func")
}

// handler interface beyond code: -> Which is Predefined!!
// type handler interface{
// 	ServeHttp( w http.ResponseWriter,r *http.Request)
// }

func main() {

	var d hotdog
	//  LISTEN AND SERVER (ADDRESS, HANDLER)
	http.ListenAndServe(":8080", d)
}
