# Data Engineering with Go (for Python practicioners)

> “Go is what Python should have been.”  
> A bold claim — worth proving with code.

This repo explores how to do everyday **data engineering tasks** in **Go**, comparing them to their Python equivalents.  

You can consider this a series of articles/tutorials to become familiar with Go for data engineering.

# Setup

```bash
mkdir go-data-engineering && cd go-data-engineering

go mod init github.com/dmschauer/go-data-engineering
```

```bash
# run directly
go run ./examples/01_csv_operations/main.go

# build and run
go build -o bin/csv_example ./examples/01_csv_operations/main.go

./bin/csv_example
```



# Shell like work
```bash
go install github.com/x-motemen/gore/cmd/gore@latest

gore
```


