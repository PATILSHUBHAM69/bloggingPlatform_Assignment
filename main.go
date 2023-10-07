package main

import (
	database "App/database"
	routes "App/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	r := gin.Default()
	database.InitDB()
	routes.SetupBlogRoutes(r)
	r.Run(":8080")
}
