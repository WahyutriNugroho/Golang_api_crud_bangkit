package main

import (
	"golang_api/controllers"
	"golang_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	//model
	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Selamat Datang di Fire Departement"})
	})

	r.GET("/prediksi", controllers.PrediksiTampil)
	r.POST("/prediksi", controllers.PrediksiTambah)
	r.PUT("/prediksi/:id", controllers.PrediksiUbah)
	r.DELETE("/prediksi/:id", controllers.PrediksiHapus)

	r.Run()
}
