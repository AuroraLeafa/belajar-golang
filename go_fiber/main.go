package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {
	app := fiber.New(
		fiber.Config{
			IdleTimeout:  time.Second * 5,
			WriteTimeout: time.Second * 5,
			ReadTimeout:  time.Second * 5,
			Prefork:      true,
		})

	// Mencoba middleware, Middleware hanya berjalan di route /api saja,
	//jika ingin berjalan disemua route maka tidak perlu menambahkan route
	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("Middleware before processing")
		err := c.Next()
		fmt.Println("Middleware after processing")
		return err
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})

	if fiber.IsChild() {
		fmt.Println("Child Proccess")
	} else {
		fmt.Println("Parent Process")
	}

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
