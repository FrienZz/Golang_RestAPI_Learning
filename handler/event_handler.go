package handler

import (
	"net/http"
	"strconv"

	"github.com/FrienZz/Golang_RestAPI_Learning/models"
	"github.com/FrienZz/Golang_RestAPI_Learning/service"
	"github.com/gin-gonic/gin"
)

type eventHandler struct {
	service service.EventService
}

func NewEventHandler(service service.EventService) *eventHandler{
	return &eventHandler{service : service}
}

func (h *eventHandler) CreateEvent(c *gin.Context){
	var newEvent models.Event
	err := c.BindJSON(&newEvent)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}
	
	err = h.service.CreateEvent(newEvent)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message" : "Event created successfully!"})
}

func (h *eventHandler) GetAllEvent(c *gin.Context){
	result,err := h.service.GetAllEvent()

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *eventHandler) GetEvent(c *gin.Context){
	id,err := strconv.Atoi(c.Param("id"))

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	result,err := h.service.GetEvent(id)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *eventHandler) UpdateEvent(c *gin.Context){
	var updateEvent models.Event
	id,err := strconv.Atoi(c.Param("id"))

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	err = c.BindJSON(&updateEvent)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}
	
	err = h.service.UpdateEvent(updateEvent.Name,updateEvent.Description,id)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message" : "Event created successfully!"})
}

func (h *eventHandler) DeleteEvent(c *gin.Context){
	id,err := strconv.Atoi(c.Param("id"))

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}
	
	err = h.service.DeleteEvent(id)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Event deleted succesfully!"})
}
