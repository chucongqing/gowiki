package main

import (
	_ "fmt"

	"github.com/gowiki/server"
)

func main() {
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
	server.Run()
}
