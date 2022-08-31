package main

import (
	"cloudflare-pgbouncer/Database"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	Database.PostgresConnection()
	defer Database.PostgresDB.Close()

	// create user

	_, queryError := Database.PostgresDB.Exec(`INSERT INTO "User" ("UserName" , "Password") VALUES ($1 , $2) `, "mimdl", "deli")
	if queryError != nil {
		fmt.Println(1)
		fmt.Println(queryError)
		return
	}

	router := gin.Default()
	router.GET("/readUser", Database.ReadUser)
	router.Run()

}
