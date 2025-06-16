package handler

import (
	"net/http"

	"github.com/FrienZz/Golang_RestAPI_Learning/internal/models"
	"github.com/FrienZz/Golang_RestAPI_Learning/internal/port"
	"github.com/gin-gonic/gin"
)

type eventHandler struct {
	service port.EventService
}

func NewEventHandler(service port.EventService) *eventHandler{
	return &eventHandler{service : service}
}

func (h *eventHandler) CreateEvent(c *gin.Context){

	var newEvent models.Event
	err := c.BindJSON(&newEvent)
	if err != nil{
		HandleError(c,err)
		return
	}
	
	err = h.service.CreateEvent(newEvent)

	if err != nil{
		HandleError(c,err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message" : "Event has been successfully created!"})
}

func (h *eventHandler) GetAllEvent(c *gin.Context){

	result,err := h.service.GetAllEvent()

	if err != nil{
		HandleError(c,err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *eventHandler) GetEvent(c *gin.Context){
	id := c.Param("id")

	result,err := h.service.GetEvent(id)

	if err != nil{
		HandleError(c,err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *eventHandler) UpdateEvent(c *gin.Context){
	var updateEvent models.Event
	id := c.Param("id")

	err := c.BindJSON(&updateEvent)

	if err != nil{
		HandleError(c,err)
		return
	}
	
	err = h.service.UpdateEvent(updateEvent.Name,updateEvent.Description,id)

	if err != nil{
		HandleError(c,err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message" : "Event has been successfully updated!"})
}

func (h *eventHandler) DeleteEvent(c *gin.Context){
	id := c.Param("id")
	
	err := h.service.DeleteEvent(id)

	if err != nil{
		HandleError(c,err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Event has been successfully deleted!"})
}
