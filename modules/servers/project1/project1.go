package project1

import (
	"HexagonalModel/modules/project1/handlers"
	"HexagonalModel/modules/project1/repositories"
	"HexagonalModel/modules/project1/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Project1(router *gin.Engine, db *sql.DB) {

	r := repositories.NewRepositorie(db)
	s := services.NewService(r)
	h := handlers.NewHandler(s)

	router.POST("/api/test", h.PostRegisterHandler)

}
