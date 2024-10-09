package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "anigo/db"
    "anigo/models"
)

func GetAnimeList(c *gin.Context) {
    var animeList []models.Anime
    db.DB.Find(&animeList)
    c.JSON(http.StatusOK, animeList)
}

func AddAnime(c *gin.Context) {
    var anime models.Anime

    if err := c.ShouldBindJSON(&anime); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.DB.Create(&anime).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create anime"})
        return
    }

    c.JSON(http.StatusCreated, anime)
}

func GetAnimeById(c *gin.Context) {
    var anime models.Anime

    if err := db.DB.First(&anime, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Anime not found"})
        return
    }

    c.JSON(http.StatusOK, anime)
}

func UpdateAnimeById(c *gin.Context) {
    var anime models.Anime

    if err := db.DB.First(&anime, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Anime not found"})
        return
    }

    if err := c.ShouldBindJSON(&anime); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.DB.Save(&anime)
    c.JSON(http.StatusOK, anime)
}

func DeleteAnimeById(c *gin.Context) {
    if err := db.DB.Delete(&models.Anime{}, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Anime not found"})
        return
    }

    c.Status(http.StatusNoContent)
}
