package main

import (
	"github.com/mpbauer/hackzurich-2018-drugify-server/config"
	"github.com/mpbauer/hackzurich-2018-drugify-server/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/mpbauer/hackzurich-2018-drugify-server/handlers"
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin/binding"
	"github.com/mpbauer/hackzurich-2018-drugify-server/validators"
)

var conf = config.Config{}
var apiHandler = handlers.Handler{}

func init() {
	conf.Read()

	db, err := models.NewDB(conf.Database.Server, conf.Database.Database)

	if err != nil {
		log.Panic(err)
	}

	apiHandler = handlers.Handler{DB: db}
	log.Println("Server: ", conf.Database.Server, " Database: ", conf.Database.Database)
}

func main() {

	router := gin.Default()

	binding.Validator = new(validators.DefaultV9Validator)

	api := router.Group("/api")
	{
		api.GET("/drugs/:swissMedicId", apiHandler.FindDrug)
		api.GET("/users/:userId/drugs", apiHandler.GetFullDrugHistoryItemsHandler)
		api.POST("/users/:userId/drugs", apiHandler.CreateDrugHistoryItemHandler)
	}

	// configure default route handler
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router.Run(":3000")

}
