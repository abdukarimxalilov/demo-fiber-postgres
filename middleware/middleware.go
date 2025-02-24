package middleware
import (
    "github.com/gofiber/fiber"
    "github.com/gofiber/basicauth"
    "github.com/abdukarimxalilov/demo-fiber-postgres/config"
)

func AuthReq() func(*fiber.Ctx) {
    cfg := basicauth.Config{
        Users: map[string]string{
          config.Config("postgres"): config.Config("your_db"),
        },
      }
    err := basicauth.New(cfg);
    return err
}