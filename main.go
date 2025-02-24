package main 

import (
"github.com/gofiber/fiber" 
"log"
"github.com/gofiber/fiber/middleware"
"github.com/abdukarimxalilov/demo-fiber-postgres/database"
"github.com/abdukarimxalilov/demo-fiber-postgres/router"

_ "github.com/lib/pq"
)

func main() { 
  if err := database.Connect(); err != nil {
      log.Fatal(err)
    }

  app := fiber.New()

  app.Use(middleware.Logger())
  router.SetupRoutes(app)

  app.Listen(3000) 
}