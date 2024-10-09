package main

import (
    "github.com/gin-gonic/gin"
    "anigo/db"
    "anigo/controllers"
)

func main() {
    db.ConnectDatabase()

    router := gin.Default()

    router.GET("/anime", controllers.GetAnimeList)
    router.POST("/anime", controllers.AddAnime)
    router.GET("/anime/:id", controllers.GetAnimeById)
    router.PUT("/anime/:id", controllers.UpdateAnimeById)
    router.DELETE("/anime/:id", controllers.DeleteAnimeById)

    router.Run(":8888")
}
