package database
import (
    "database/sql"
    "fmt"
    "strconv"
    "github.com/abdukarimxalilov/demo-fiber-postgres/config" 
)

var DB *sql.DB

func Connect() error {
    var err error
    p := config.Config("5432")
    port, err := strconv.ParseUint(p, 10, 32) 
    if err != nil {
        fmt.Println("Error parsing str to int")
    }
    DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("localhost"), port, config.Config("postgres"), config.Config("your_db_pass"), config.Config("postgres")))

    if err != nil {
        return err
    }
    if err = DB.Ping(); err != nil {
        return err
    }
    CreateProductTable()
    fmt.Println("Connection Opened to Database")
    return nil
}