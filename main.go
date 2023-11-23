package main 

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Todos struct {
	title string 
	todo string 
	do bool 
}

func main() {
	var list []Todos;
	list = append(list , Todos{title : "test" , todo :"test",do : false})

	app:= fiber.New();

	fmt.Println("The server is running on port 3200");
	app.Get("/" , func(c *fiber.Ctx) error {
	  return c.SendString("Home page");
	});
	app.Listen(":3200");

}
