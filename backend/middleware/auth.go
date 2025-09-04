package middleware

import (
	"split-bill/backend/config"
	"split-bill/backend/repository"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthOptions struct {
	JwtConfig      *config.JwtConfig
	CookieName     string
	UserRepository repository.UserRepository
}

// NewAuth returns a Fiber handler that enforces JWT on protected routes.
func NewAuth(jwtConfig *config.JwtConfig, userRepository repository.UserRepository) fiber.Handler {
	// sensible defaults
	opt := &AuthOptions{
		CookieName:     "token",
		JwtConfig:      jwtConfig,
		UserRepository: userRepository,
	}

	return func(c *fiber.Ctx) error {
		raw := ""

		if v := c.Cookies(opt.CookieName); v != "" {
			raw = v
		}

		if raw == "" {
			return fiber.ErrUnauthorized
		}

		claims, err := opt.JwtConfig.ValidateToken(raw)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		if exp := claims.Exp; exp < time.Now().Unix() {
			return fiber.ErrUnauthorized
		}

		user, err := opt.UserRepository.FindByUUID(claims.UserID)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		c.Locals("user", user)
		return c.Next()
	}
}
