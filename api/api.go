// api/api.go
package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"user_management_service/database"
)

type API struct {
	db     *database.Database
	router *gin.Engine
}

func NewAPI(db *sql.DB, router *gin.Engine) *API {
	return &API{db: database.NewDatabase(db), router: router}
}

func (api *API) SetupRoutes() {
	api.router.POST("/users", api.createUser)
	api.router.GET("/users/:id", api.getUser)
	api.router.PUT("/users/:id", api.updateUser)
	api.router.DELETE("/users/:id", api.deleteUser)
}
