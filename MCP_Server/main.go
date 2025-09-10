package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	fmt.Println("MCP Server-y sksum e ashxatanqy...")

	connStr := "server=localhost\\SQLEXPRESS;database=DB;trusted_connection=yes"

	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		fmt.Println("Sxal kapakcum:", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Ping sxal:", err)
		return
	}

	fmt.Println("db-in normal kapvecinq:")
}
