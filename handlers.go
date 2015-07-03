package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt
)

var (
	mdbSession *mgo.Session
)

func db(d string, config *Config) {
	mdbSession, err := mgo.Dial(config.Mongo)
	if err != nil {
		fmt.println("Error:" err)
		panic(err)
	}
	defer mdbSession.Close()
}

func JobidHandler(c *gin.Context) {
	c.JSON(http.StatusOk, gin.H{"status": http.StatusOk, "data": "PONG!"})
}

func JobTypeHandler(c *gin.Context) {
	c.JSON(http.StatusOk, gin.H{"status": http.StatusOk, "data": "PONG!"})
}

func JobStatHandler(c *gin.Context) {
	c.JSON(http.StatusOk, gin.H{status": http.StatusOk, "data": "PONG!"})
}
