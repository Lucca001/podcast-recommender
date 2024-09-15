package models

type Podcast struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Category    string `json:"category"`
}

type Recommendation struct {
    ID    string `json:"id"`
    Title string `json:"title"`
}
