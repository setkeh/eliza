package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// REST Server
	server := gin.Default()

	// Routes

	// #JobID
	server.GET("/jobID/:id", JobidHandler)
	server.POST("/jobID/:id", JobidHandler)

	// Job Type
	server.GET("/JobType", JobTypeHandler)
	server.POST("/JobType", JobTypeHandler)

	// Job Status
	server.GET("/JobStatus", JobStatHandler)
	server.POST("/JobStatus", JobStatHandler)

	// Start Server
	server.Run(":80")
}
