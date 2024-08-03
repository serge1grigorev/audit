package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "strconv"
)

type log struct {
    ID int `json:"id"`
    Message string `json:"message"`
}

var logs = []log {
    {ID: 1, Message: "First message"},
    {ID: 2, Message: "Second message"},
}

func getLogs(context *gin.Context) {
    context.IndentedJSON(http.StatusOK, logs)
}

func getLogById(context *gin.Context) {
    var idStr = context.Param("id")
    var id, err = strconv.Atoi(idStr)
    if err != nil {
        context.IndentedJSON(http.StatusBadRequest, gin.H{"message": idStr + " is not a number"})
        return
    }

    for _, log := range logs {
        if log.ID == id {
            context.IndentedJSON(http.StatusOK, log)
            return
        }
    }
    context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album with id " + idStr + " not found"})
}

func addLog(context *gin.Context) {
    var newLog log

    var err = context.BindJSON(&newLog)
    if err != nil {
        return
    }

    logs = append(logs, newLog)
    context.IndentedJSON(http.StatusCreated, newLog)
}

func main() {
    var router = gin.Default()
    router.GET("/logs", getLogs)
    router.GET("/logs/:id", getLogById)
    router.POST("/add-log", addLog)

    router.Run(":8082")
}