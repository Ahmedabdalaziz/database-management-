package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

type User struct {
	ID        int
	Name      string
	Email     string
	City      string
	CreatedAt string
}

func main() {
	user := "root"
	pass := "rootroot"
	host := "127.0.0.1"
	port := "3306"
	dbname := "golang"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("error opening db connection:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln("error pinging db:", err)
	}

	fmt.Println("connected to database")

	users := getAllUsers()
	fmt.Println(users)
}

func insertUser(name, email, city string) int {
	res, err := db.Exec("insert into users (name, email, city) values (?, ?, ?)", name, email, city)
	if err != nil {
		log.Fatalln("error inserting user:", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalln("error getting inserted user id:", err)
	}
	return int(id)
}

func updateUser(id int, name, email, city string) {
	_, err := db.Exec("update users set name = ?, email = ?, city = ? where id = ?", name, email, city, id)
	if err != nil {
		log.Fatalln("error updating user:", err)
	}
	fmt.Println("updated user with id", id)
}

func deleteUser(id int) {
	_, err := db.Exec("delete from users where id = ?", id)
	if err != nil {
		log.Fatalln("error deleting user:", err)
	}
	fmt.Println("deleted user with id", id)
}

func getUserByID(id int) *User {
	row := db.QueryRow("select id, name, email, city, created_at from users where id = ?", id)

	var u User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.City, &u.CreatedAt)
	if err == sql.ErrNoRows {
		fmt.Println("no user found with that id")
		return nil
	} else if err != nil {
		log.Fatalln("error fetching user:", err)
	}
	return &u
}

func getAllUsers() []User {
	rows, err := db.Query("select id, name, email, city, created_at from users order by created_at desc")
	if err != nil {
		log.Fatalln("error fetching users:", err)
	}
	defer rows.Close()

	var list []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.City, &u.CreatedAt)
		if err != nil {
			log.Fatalln("error scanning user row:", err)
		}
		list = append(list, u)
	}
	return list
}

func findUsersByCity(city string) []User {
	rows, err := db.Query("select id, name, email, city, created_at from users where city = ?", city)
	if err != nil {
		log.Fatalln("error searching users by city:", err)
	}
	defer rows.Close()

	var list []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.City, &u.CreatedAt)
		if err != nil {
			log.Fatalln("error reading user row:", err)
		}
		list = append(list, u)
	}
	return list
}

func countUsers() int {
	var count int
	err := db.QueryRow("select count(*) from users").Scan(&count)
	if err != nil {
		log.Fatalln("error counting users:", err)
	}
	return count
}
