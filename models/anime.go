package models

import (
    "gorm.io/gorm"
)

type Anime struct {
    gorm.Model
    Title       string  `json:"title"`
    Description string  `json:"description"`
    Genre       string  `json:"genre"`
    Rating      float64 `json:"rating"`
}
