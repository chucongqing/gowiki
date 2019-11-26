package main

import (
	_ "fmt"

	// "rsc.io/quote"
	_ "github.com/gowiki/server"
	// "github.com/gowiki/echoserver"

	"net/http"
	"github.com/labstack/echo/v4"
)

func runServer(){
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":4399"))
}

func main() {
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
	// quote.Hello()
	runServer();
	// server.Run()
}

