package utils

import "github.com/gofiber/fiber/v2"

// StaticKey returns a function that always returns the same cache key
func StaticKey(key string) func(c *fiber.Ctx) string {
	return func(c *fiber.Ctx) string {
		return key
	}
}

// Example dynamic key helper for pagination
func PaginatedKey(base string) func(c *fiber.Ctx) string {
	return func(c *fiber.Ctx) string {
		page := c.Query("page", "1")
		return base + "_page_" + page
	}
}
