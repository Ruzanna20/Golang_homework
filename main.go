package main

import (
	"database/sql"
	"fmt"
	"log"
	"mcp_sql_server_/repository"
	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	fmt.Println("MCP Server starts working...")

	server := "localhost"
	port := 1433
	user := "SA"
	password := "Ruzanna04!"
	database := "TravelDB"

	connStr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
	server,user,password,port,database)

	db,err := sql.Open("sqlserver",connStr)
	if err != nil {
		log.Fatal("Incorrect connection:",err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Ping error:",err)
	}
	fmt.Println("Connected to DB successfully.")


	results,err := repository.GetContactOfUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	for _,res := range results {
			fmt.Println("Phone of User:",res["FirstName"],"-",res["Phone"])		
		}
	
	fmt.Println("")

	users,err := repository.GetAllUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	for _,u := range users {
		fmt.Print("User:",u,"\n")
	}

	}