package main

import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type IRIS struct{
	Sepal_Length float64
	Sepal_Width float64
	Petal_Length float64
	Petal_Width float64
	Species string
}

func main(){
	Postgres_connect()
}

func Postgres_connect(){
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=****** dbname=postgres sslmode=disable")
	defer db.Close()
	if err != nil{
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT * FROM iris")

    if err != nil {
        fmt.Println(err)
	}
	
	var iris []IRIS
    for rows.Next() {
        var e IRIS
        rows.Scan(&e.Sepal_Length, &e.Sepal_Width, &e.Petal_Length, &e.Petal_Width, &e.Species)
        iris = append(iris, e)
    }
    fmt.Printf("%v", iris)
}

