package main

import (
  "fmt"
  "github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New();
  fmt.Println("lanci hamid");
  app.Listen(":3030");
}
