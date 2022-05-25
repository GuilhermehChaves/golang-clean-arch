package main

import (
	"cleanarch/src/external/web/route"
	"log"
)

func init() {
	log.Println("Initing...")
}

func main() {
	ginRoute := route.NewGinRoute().GetEngine()
	ginRoute.Run(":8080")
}
