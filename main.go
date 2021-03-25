package main

import (
	"fmt"

	"github.com/MarekWojt/GoMan/http"
	"github.com/MarekWojt/GoMan/orm"
)

func run() {
	fmt.Println("starting http server")
	err := http.Run()
	fmt.Println(err)
}

func main() {
	orm.Init()
	println("orm initialized")
	run()
}
