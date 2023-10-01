package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sriharshamur/lenslocked/models"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (pgc PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", pgc.Host, pgc.Port, pgc.User, pgc.Password, pgc.Database, pgc.SSLMode)
}

func main() {
	pgConfig := models.DefaultPostgresConfig()
	db, err := models.Open(pgConfig)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")

	us := models.UserService{
		DB: db,
	}
	user, err := us.Create("bob1@test.com", "bobspassword")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	// _, err = db.Exec(`
	// 	CREATE TABLE IF NOT EXISTS users (
	// 		id SERIAL PRIMARY KEY,
	// 		name TEXT,
	// 		email TEXT UNIQUE NOT NULL
	// 	);

	// 	CREATE TABLE IF NOT EXISTS orders (
	// 		id SERIAL PRIMARY KEY,
	// 		user_id INT NOT NULL,
	// 		amount INT,
	// 		description TEXT
	// 	);
	// `)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Tables Created!")

	// Create user
	// name := "Sri Harsh Amur"
	// email := "sriharshamur1@gmail.com"
	// _, err = db.Exec(`
	// 	INSERT INTO users (name, email)
	// 	VALUES ($1, $2);`, name, email)
	// row := db.QueryRow(`
	// 	INSERT INTO users (name, email)
	// 	VALUES ($1, $2) RETURNING id;
	// `, name, email)
	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User Created with id =", id)

	// Get User Information
	// id := 1
	// row := db.QueryRow(`SELECT name, email
	// FROM users
	// WHERE id=$1`, id)
	// var name, email string
	// err = row.Scan(&name, &email)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("User Information: name=%s, email=%s\n", name, email)

	// Create orders for a user
	// userId := 1
	// for i := 1; i <= 5; i++ {
	// 	amount := i * 100
	// 	description := fmt.Sprintf("Fake Order #%d", i)
	// 	_, err = db.Exec(`
	// 		INSERT INTO orders (user_id, amount, description)
	// 		VALUES ($1 ,$2, $3)
	// 	`, userId, amount, description)

	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("Created 5 Orders")

	// type Order struct {
	// 	ID          int
	// 	UserID      int
	// 	Amount      int
	// 	Description string
	// }
	// var orders []Order
	// userId := 1
	// rows, err := db.Query(`
	// 	SELECT id, amount, description
	// 	FROM orders
	// 	WHERE user_id=$1
	// `, userId)
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var order Order
	// 	order.UserID = userId
	// 	err = rows.Scan(&order.ID, &order.Amount, &order.Description)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	orders = append(orders, order)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Orders:", orders)
}
