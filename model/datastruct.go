package model

type Cake struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Image       *string `json:"image"`
	Created_at  string  `json:"created_at"`
	Updated_at  *string `json:"updated_at"`
}

type MessageData struct {
	Status  bool
	Message string
}

type CreateCake struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Rating      float32 `json:"rating" binding:"required"`
	Image       *string `json:"image" binding:"required"`
}

type UpdateCake struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Rating      float32 `json:"rating" binding:"required"`
	Image       *string `json:"image" binding:"required"`
}
