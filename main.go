package main

import (
	router "github.com/faabioms/ping-pong-throttling/router"
)

func main() {
	r := router.Router()

	r.Run(":8080")
}