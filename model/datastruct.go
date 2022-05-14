package model

type Cake struct {
	Id          string  `json:"id" example:"1"`
	Title       string  `json:"title" example:"tahu bulat"`
	Description string  `json:"description" example:"tahu bulat digoreng dadakan"`
	Rating      float32 `json:"rating" example:"9"`
	Image       *string `json:"image" example:"http://linkketahubulat.jpg"`
	Created_at  string  `json:"created_at" example:"false"`
	Updated_at  *string `json:"updated_at" example:"2022-05-14 22:58:18"`
}

type MessageData struct {
	Status  bool
	Message string
}

type CreateCake struct {
	Title       string  `json:"title" binding:"required" example:"tahu bulat"`
	Description string  `json:"description" binding:"required" example:"tahu bulat digoreng dadakan"`
	Rating      float32 `json:"rating" binding:"required" example:"9"`
	Image       *string `json:"image" binding:"required" example:"http://linkketahubulat.jpg"`
}

type UpdateCake struct {
	Title       string  `json:"title" binding:"required" example:"tahu bulat"`
	Description string  `json:"description" binding:"required" example:"tahu bulat digoreng dadakan"`
	Rating      float32 `json:"rating" binding:"required" example:"9"`
	Image       *string `json:"image" binding:"required" example:"http://linkketahubulat.jpg"`
}
