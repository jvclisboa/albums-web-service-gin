package main

import (
   "example/web-service-gin/router"
)

func main() {
    router := router.InitRouter()

    router.Run("localhost:8080")
}