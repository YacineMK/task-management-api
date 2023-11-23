package main 

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Todos struct {
	id int
	title string 
	todo string 
	do bool 
}

func main() {
	var list []Todos;
	list = append(list , Todos{id :1 ,title : "test" , todo :"test",do : false})

	app:= fiber.New();
	fmt.Println("The server is running on port 3200");

	app.Post("/addtodo",func(c *fiber.Ctx) error {
		var todo Todos ;
		var i int ;	
		func check() string{
			for i range list {
				if (todo.title == list[i].title && list[i].title == false) {
					return "you have allredy this todo ";
				}
			}
		}
		list = append();
	})

	app.Get("/" , func(c *fiber.Ctx) error {
	  return c.SendString("Home page");
	});
	app.Listen(":3200");

}
