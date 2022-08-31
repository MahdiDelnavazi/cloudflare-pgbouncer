package Database

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// attach chunk to order
func ReadUser(context *gin.Context) {

	// read order
	var user []UserEntity
	queryError := PostgresDB.Select(&user, `SELECT "Id" FROM "User"`)
	if queryError != nil {
		log.Println(queryError)
		context.JSON(http.StatusInternalServerError, gin.H{"message": queryError})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "successful", "user id": user})

}
