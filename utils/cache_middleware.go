package utils

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

// CacheMiddleware caches GET request responses in Redis
// keyFunc generates a cache key per request
func CacheMiddleware(ttl time.Duration, keyFunc func(c *fiber.Ctx) string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() != fiber.MethodGet {
			return c.Next()
		}

		key := keyFunc(c)

		// Try cache
		cached, err := GetCache(key)
		if err == nil {
			c.Set("X-Cache", "HIT")
			return c.SendString(cached)
		}

		c.Set("X-Cache", "MISS")

		// Let controller run
		if err := c.Next(); err != nil {
			return err
		}

		// Store response in Redis
		body := c.Response().Body()
		SetCache(key, string(body), ttl)

		return nil
	}
}
