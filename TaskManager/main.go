package main

import (
	"TaskManager/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080") 
}
