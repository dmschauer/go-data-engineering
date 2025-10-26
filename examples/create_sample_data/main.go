package main

import (
    "fmt"
    "os"

    "github.com/go-gota/gota/dataframe"
)

type User struct {
    Name     string
    Age      int
    Accuracy float64
    ignored  bool
}

func main() {
    users := []User{
        {"Aram", 17, 0.2, true},
        {"Juan", 18, 0.8, true},
        {"Ana", 22, 0.5, true},
    }

    df := dataframe.LoadStructs(users)

    outputPath := "../../data/input/users.csv"

    f, err := os.Create(outputPath)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    if err := df.WriteCSV(f); err != nil {
        panic(err)
    }

    fmt.Printf("CSV written to %s\n", outputPath)
}
