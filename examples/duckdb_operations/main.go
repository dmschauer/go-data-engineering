package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/duckdb/duckdb-go/v2"
)

func main() {
    db, err := sql.Open("duckdb", "")
    if err != nil {
        log.Fatalf("failed to open DuckDB: %v", err)
    }
    defer db.Close()

    query := `
        COPY (
            SELECT Name, Age, Accuracy
            FROM read_csv_auto('data/input/users.csv')
            WHERE Age > 18
        ) TO 'data/output/result.csv' (FORMAT CSV);
    `

	db.Exec(query)

    fmt.Println("✅ Filtered result written")
}
