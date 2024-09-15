package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "podcast-recommender/services"
)

func GetRecommendations(c *gin.Context) {
    userID := c.Query("user_id")
    recommendations, err := services.GetRecommendationsForUser(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, recommendations)
}
