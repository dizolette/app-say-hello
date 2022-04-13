package main

import (
	"fmt"

	go_say_hello "github.com/dizolette/go-say-hello"
)

/*

cara mendapatkan module, harus menggunakan command berikut :
go get <nama-module>

di go.mod akan menambahkan require module yang di get

*/
func main() {
	hello := go_say_hello.SayHello()
	fmt.Println(hello)
}
