package main

import (
	"go-gin-web-demo/app/river-api/internal/conf"
)

func main() {
	println("hello")
	router := conf.SetupRouter()
	conf.SetupDB()

	_ = router.Run()

}
