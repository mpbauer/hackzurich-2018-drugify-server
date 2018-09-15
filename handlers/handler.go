package handlers

import (
	"github.com/mpbauer/hackzurich-2018-drugify-server/models"
	"github.com/mpbauer/hackzurich-2018-drugify-server/errorhandling"
	"github.com/gin-gonic/gin"
	"net/http"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

type Handler struct {
	DB *models.DB
}

func (h *Handler) CreateDrugHistoryItemHandler(c *gin.Context) {
	var drugHistory models.DrugHistory

	if err := c.BindJSON(&drugHistory); err != nil {
		log.WithError(err).Warn("Invalid payload request")
		return
	}
	drugHistory.ID = bson.NewObjectId().Hex()

	if err := h.DB.InsertDrugHistoryItem(drugHistory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusCreated, drugHistory)
}

func (h *Handler) GetFullDrugHistoryItemsHandler(c *gin.Context) {
	username := c.Param("userId")

	drugHistoryItems, err := h.DB.GetFullDrugHistory(username)

	if err != nil {
		log.WithError(err).Error("An unexpected error occured")
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, drugHistoryItems)
}


func (h *Handler) FindDrug(c *gin.Context) {
	swissMedicId := c.Param("swissMedicId")

	drug, err := h.DB.FindDrug(swissMedicId)

	if err != nil {
		switch err.(type) {
		case *errorhandling.ErrNotFound:
			log.WithFields(log.Fields{
				"swiss medic id:": swissMedicId,
			}).Info("No project found")
			c.JSON(http.StatusNotFound, gin.H{})
			return
		default:
			log.WithError(err).Error("An unexpected error occured")
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
	}

	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, drug)
}

func (h *Handler) HasConflict(c *gin.Context) {
	swissMedicId := c.Param("swissMedicId")

	if(swissMedicId == "40770"){
		c.JSON(http.StatusConflict, gin.H{})
	}else{
		c.JSON(http.StatusOK, gin.H{})
	}
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Header("Access-Control-Allow-Origin", "*")
}