package routes

import(
	"net/http"
	"strconv"

	"github.com/gevorgalaverdyan/go-playground/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "internal err"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context){
	var newEvent models.Event

	err := ctx.ShouldBindJSON(&newEvent)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message":"binding issue"})
		return 
	}

	errMsg := newEvent.Save()
	if errMsg.Message != "" {
		ctx.JSON(http.StatusInternalServerError, errMsg.Message)
		return
	}
	ctx.JSON(http.StatusCreated, newEvent)
}

func getEvent(ctx *gin.Context){
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}

	event, customErr := models.GetById(id)

	if customErr.Message != "" {
		ctx.JSON(http.StatusInternalServerError, customErr.Message)
		return 
	}

	ctx.JSON(http.StatusOK, event)
}

func updateEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}

	_, customErr := models.GetById(id)
	if customErr.Message != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message":"exists issue"})
		return 
	}

	var e models.Event
	err = ctx.ShouldBindJSON(&e)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message":"binding issue"})
		return 
	}

	e.ID = id
	e, customErr = e.Update()

	if customErr.Message != "" {
		ctx.JSON(http.StatusInternalServerError, customErr.Message)
		return 
	}

	ctx.JSON(http.StatusOK, e)
}

func deleteEvent(ctx *gin.Context){
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}

	e, customErr := models.GetById(id)
	if customErr.Message != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message":"exists issue"})
		return 
	}

	customErr = e.Delete()

	if customErr.Message != "" {
		ctx.JSON(http.StatusInternalServerError, customErr.Message)
		return 
	}

	ctx.JSON(http.StatusOK, e)
}