package main

import (
  "fmt"
  "github.com/gofiber/fiber/v2"
)

func main() {
   app:= fiber.New();
    fmt.Println("the server run in port 3200")
    app.Listen(":3200")
}
