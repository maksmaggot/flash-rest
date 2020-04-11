package controllers

import (
	"flash-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func FindBooks(context *gin.Context) {
	db := context.MustGet("db").(*gorm.DB)

	var books []models.Book
	db.Find(&books)

	context.JSON(http.StatusOK, gin.H{"data": books})
}

func FindBook(context *gin.Context) {
	db := context.MustGet("db").(*gorm.DB)

	var book models.Book
	if err := db.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record Not Found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": book})
}

func CreateBook(context *gin.Context) {
	db := context.MustGet("db").(*gorm.DB)

	var input CreateBookInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	db.Create(&book)

	context.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(context *gin.Context) {
	db := context.MustGet("db").(*gorm.DB)

	var book models.Book
	if err := db.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateBookInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&book).Updates(input)

	context.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(context *gin.Context) {
	db := context.MustGet("db").(*gorm.DB)

	var book models.Book
	if err := db.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Not Found!"})
		return
	}

	db.Delete(&book)

	context.JSON(http.StatusOK, gin.H{"data": true})
}
