package auth

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

const JwtSecret = "asecret"

func RequiredAuth() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SuccessHandler: func(c *fiber.Ctx) error {
			err := c.Status(200).JSON(fiber.Map{
				"Authorization": "Success",
			})
			if err != nil {
				return err
			}
			return c.Next()
		},
		SigningKey: []byte(JwtSecret),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			err = ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"errorCode": err.Error(),
				"error":     "Unauthorized",
			})
			if err != nil {
				return err
			}
			return nil
		},
	})
}
