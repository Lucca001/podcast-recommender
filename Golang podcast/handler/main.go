package main

import (
    "github.com/gin-gonic/gin"
    "podcast-recommender/handlers"
)

func main() {
    r := gin.Default()

    r.GET("/podcasts", handlers.GetPodcasts)
    r.GET("/recommendations", handlers.GetRecommendations)

    r.Run(":8080")
}
