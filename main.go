package main 

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)



func main() {
	app:= fiber.New();
	fmt.Println("The server is running on port 3200");
	app.Get("/" , func(c *fiber.Ctx) error {
	  return c.SendString("Home page");
	});
	app.Listen(":3200");

}
