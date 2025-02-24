package handler

import (
    "log"
    "database/sql"
    "github.com/gofiber/fiber"
    "github.com/abdukarimxalilov/demo-fiber-postgres/model"
    "github.com/abdukarimxalilov/demo-fiber-postgres/database"
)

func GetAllProducts(c *fiber.Ctx) {
    rows, err := database.DB.Query("SELECT name, description, category, amount FROM products order by name")
    if err != nil {
        c.Status(500).JSON(&fiber.Map{
            "success": false,
            "error": err,
          })
        return
    }
    defer rows.Close()
    result := model.Products{}
    for rows.Next() {
        product := model.Product{}
        err := rows.Scan(&product.Name, &product.Description, &product.Category, &product.Amount)
        if err != nil {
            c.Status(500).JSON(&fiber.Map{
                "success": false,
                "error": err,
              })
            return
        }
        result.Products = append(result.Products, product)
    }
    if err := c.JSON(&fiber.Map{
        "success": true,
        "product":  result,
        "message": "All product returned successfully",
      }); err != nil {
        c.Status(500).JSON(&fiber.Map{
            "success": false,
            "message": err,
          })
        return
    }
}

func GetSingleProduct(c *fiber.Ctx) {
    id := c.Params("id")
    product := model.Product{}
    row, err := database.DB.Query("SELECT * FROM products WHERE id = $1", id)
    if err != nil {
        c.Status(500).JSON(&fiber.Map{
            "success": false,
            "message": err,
          })
        return
    }
    defer row.Close()
    for row.Next() {
    switch err := row.Scan(&id, &product.Amount, &product.Name, &product.Description, &product.Category ); err {
        case sql.ErrNoRows:
              log.Println("No rows were returned!")
              c.Status(500).JSON(&fiber.Map{
                "success": false,
                "message": err,
              })
        case nil:
            log.Println(product.Name, product.Description, product.Category, product.Amount)
        default:
              c.Status(500).JSON(&fiber.Map{
                "success": false,
                "message": err,
              })
    }
}

    if err := c.JSON(&fiber.Map{
        "success": false,
        "message": "Successfully fetched product",
        "product": product,
      }); err != nil {
        c.Status(500).JSON(&fiber.Map{
            "success": false,
            "message":  err,
          })
        return
    }

}

func CreateProduct(c *fiber.Ctx) {
    p := new(model.Product)
    if err := c.BodyParser(p); err != nil {
        log.Println(err)
        c.Status(400).JSON(&fiber.Map{
            "success": false,
            "message": err,
          })
        return
    }
    res, err := database.DB.Query("INSERT INTO products (name, description, category, amount) VALUES ($1, $2, $3, $4)" , p.Name, p.Description, p.Category, p.Amount )
    if err != nil {
        c.Status(500).JSON(&fiber.Map{
            "success": false,
            "message": err,
          })
        return
    }
    log.Println(res)

    if err := c.JSON(&fiber.Map{
        "success": true,
        "message": "Product successfully created",
        "product": p,
      }); err != nil {
        c.Status(500).JSON(&fiber.Map{
            "success": false,
            "message":  "Error creating product",
          })
        return
    }
}

func DeleteProduct(c *fiber.Ctx) {
        id := c.Params("id")
        res, err := database.DB.Query("DELETE FROM products WHERE id = $1", id)
        if err != nil {
            c.Status(500).JSON(&fiber.Map{
                "success": false,
                "error": err,
              })
            return
        }
        log.Println(res)
        if err := c.JSON(&fiber.Map{
            "success": true,
            "message": "product deleted successfully",
          }); err != nil {
            c.Status(500).JSON(&fiber.Map{
                "success": false,
                "error": err,
              })
            return
        }
}