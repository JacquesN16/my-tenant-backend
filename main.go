package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Tenant struct {
	FirstName string  `json:"firstName,string"`
	LastName  string  `json:"lastName,string"`
	Email     string  `json:"email,string"`
	StartDate int64   `json:"startDate,int64"`
	Rent      float64 `json:"rent,float64"`
	Charge    float64 `json:"charge,float64"`
}

func getUsers(c *gin.Context) {
	//Create a db connection
	db, err := sql.Open("sqlite3", "../../sqliteDb/tenant.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tenant")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var tenants []Tenant
	for rows.Next() {
		var tenant Tenant
		err = rows.Scan(&tenant.FirstName, &tenant.LastName, &tenant.Email, &tenant.StartDate, &tenant.Rent, &tenant.Charge)
		if err != nil {
			panic(err)
		}
		tenants = append(tenants, tenant)
	}
	fmt.Println("tenants", tenants)
	c.JSON(200, tenants)
}

func main() {
	r := gin.New()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, map[string]string{"message": "Hello World"})
	})

	r.GET("/users", getUsers)

	r.Run(":9000") // listen and serve on 0.0.0.0:9000
}
