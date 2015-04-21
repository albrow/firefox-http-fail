package main

import (
	"net/http"
)

func main() {
	res, err := http.Get("http://localhost:3000/todos")
	if err != nil {
		panic(err)
	}
	print(res)
}
