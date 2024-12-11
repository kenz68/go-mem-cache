package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//cache := cache.New(10*time.Minute, 20*time.Minute)

	app.Get("/set", func(c *fiber.Ctx) error {
		//cache.Set("key", "value", cache.DefaultExpiration)
		return c.SendString("Hello, World!")
	})

	//app.Get("/posts/:id", middleware.CacheMiddleware(cache),   routes.GetPosts) //commenting this route just to test the "/" endpoint.
	app.Listen(":8080")
}
