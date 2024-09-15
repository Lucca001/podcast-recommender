package services

import (
    "podcast-recommender/models"
)

func FetchPodcasts() ([]models.Podcast, error) {
    // Mocking podcast data
    podcasts := []models.Podcast{
        {ID: "1", Title: "Go Programming", Description: "Learn Go programming", Category: "Technology"},
        {ID: "2", Title: "Tech News", Description: "Latest in technology", Category: "News"},
    }
    return podcasts, nil
}

func GetRecommendationsForUser(userID string) ([]models.Recommendation, error) {
    // Mocking recommendations
    recommendations := []models.Recommendation{
        {ID: "1", Title: "Go Programming"},
        {ID: "2", Title: "Tech News"},
    }
    return recommendations, nil
}
