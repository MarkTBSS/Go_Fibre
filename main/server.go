package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	app.Listen(":3000")
}

// ==== Basic Routing =================================
/* app.Get("/", func(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
})

// Parameters
// GET http://localhost:3000/hello%20world
app.Get("/:value", func(c *fiber.Ctx) error {
	return c.SendString("value: " + c.Params("value"))
	// => Get request with value: hello world
})

// Optional parameter
// GET http://localhost:3000/john
app.Get("/:name?", func(c *fiber.Ctx) error {
	if c.Params("name") != "" {
		return c.SendString("Hello " + c.Params("name"))
		// => Hello john
	}
	return c.SendString("Where is john?")
})

// Wildcards
// GET http://localhost:3000/api/user/john
app.Get("/api/*", func(c *fiber.Ctx) error {
	return c.SendString("API path: " + c.Params("*"))
	// => API path: user/john
}) */
// =====================================================
