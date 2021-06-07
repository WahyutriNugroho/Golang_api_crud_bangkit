package controllers

import (
	"golang_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PrediksiInput struct {
	Id          int     `json:"id"`
	Class_name  string  `json:"class_name"`
	ImgUrl      string  `json:"imgUrl"`
	Date        string  `json:"date"`
	Latitude    string  `json:"latitude"`
	Longtitude  string  `json:"longtitude"`
	Probability float64 `json:"probability"`
}

//Tampil data prediksi
func PrediksiTampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var predict []models.Predict
	db.Find(&predict)
	c.JSON(http.StatusOK, gin.H{"data": predict})
}

// Tambah data prediksi
func PrediksiTambah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi input/masukkan
	var dataInput PrediksiInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses input

	predict := models.Predict{
		Class_name:  dataInput.Class_name,
		Date:        dataInput.Date,
		ImgUrl:      dataInput.ImgUrl,
		Latitude:    dataInput.Latitude,
		Longtitude:  dataInput.Longtitude,
		Probability: dataInput.Probability,
	}

	db.Create(&predict)

	c.JSON(http.StatusOK, gin.H{"data": predict})
}

// Ubah data prediksi
func PrediksiUbah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek dulu datanya
	var predict models.Predict
	if err := db.Where("id = ?", c.Param("id")).First(&predict).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data predksi tidak ditemukan"})
		return
	}

	//validasi input/masukkan
	var dataInput PrediksiInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses ubah data
	db.Model(&predict).Update(dataInput)

	c.JSON(http.StatusOK, gin.H{"data": predict})
}

// Hapus data mahasiswa
func PrediksiHapus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek dulu datanya
	var predict models.Predict
	if err := db.Where("id = ?", c.Param("id")).First(&predict).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data prediksi tidak ditemukan"})
		return
	}

	//proses hapus data
	db.Delete(&predict)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
