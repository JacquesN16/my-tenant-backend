package main

import (
	"database/sql"
	"my-tenant-backend/v2/models"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

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

	var tenants []models.Tenant
	for rows.Next() {
		var tenant models.Tenant
		err = rows.Scan(&tenant.FirstName, &tenant.LastName, &tenant.Email, &tenant.StartDate, &tenant.Rent, &tenant.Charge)
		if err != nil {
			panic(err)
		}
		tenants = append(tenants, tenant)
	}

	c.JSON(200, tenants)
}

func main() {
	r := gin.New()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, map[string]string{"message": "Hello World"})
	})

	r.GET("/users", getUsers)

	r.Run(":9223") // listen and serve on 0.0.0.0:9000
}
