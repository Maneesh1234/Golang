//cd golang-mysql
// To get the driver just run the following command:(For Mysql)
//go get -u github.com/go-sql-driver/mysql
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// // (1) CHECK THE CONNECTION
// func main() {

// 	db, err := sql.Open("mysql", "root:@/peoplelist")

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()
// }

// // (2) INSERT VALUE IN DATABASE
// func main() {

// 	db, err := sql.Open("mysql", "root:@/peoplelist")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()
// 	insert, err := db.Query("INSERT INTO peoplelists(LastName,FirstName) VALUES('Kumar','maneesh')")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer insert.Close()
// }

//(3) READ DATA
type peoplelist struct {
	PersonId  int
	FirstName string
}

// var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "root:@/peoplelist")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	peoplelists, err := db.Query("SELECT PersonId,FirstName FROM posts")
	if err != nil {
		panic(err.Error())
	}
	for peoplelists.Next() {
		var people peoplelist
		err := peoplelists.Scan(&people.PersonId, &people.FirstName)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(people)
	}
}
