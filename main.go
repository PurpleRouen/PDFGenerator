package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	r := gin.Default()
	r.POST("/generate-pdf", func(context *gin.Context) {
		var payload InscriptionPayload
		err = context.BindJSON(&payload) // the payload can be customized in watermarkData.go
		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var generatedPath *string
		generatedPath, err = generatePdf(&payload)
		if err != nil {
			context.JSON(500, gin.H{"error": err.Error()})
			return
		}

		context.File(*generatedPath)
	})

	err = r.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
