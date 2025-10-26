// examples/01_csv_operations/main.go
package main

import (
    "github.com/go-gota/gota/dataframe"
    "os"
	"fmt"
)

func main() {
    // Read CSV
    f, _ := os.Open("../../data/input/users.csv")
    df := dataframe.ReadCSV(f)
	fmt.Println(df)
    
    // Transform
	filtered := df.Filter(
		dataframe.F{
			Colname:    "Age",
			Comparator: ">",
			Comparando: 18,
		},
	)    
    // Write
    out, _ := os.Create("../../data/output/result.csv")
    filtered.WriteCSV(out)
}