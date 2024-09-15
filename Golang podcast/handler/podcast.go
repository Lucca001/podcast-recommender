package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "podcast-recommender/services"
)

func GetPodcasts(c *gin.Context) {
    podcasts, err := services.FetchPodcasts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, podcasts)
}
