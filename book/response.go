package book

type BookResponse struct {
	ID          int    `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Price       int    `json:"price" binding:"required,number"`
	Description string `json:"description" binding:"required"`
	Rating      int    `json:"rating" binding:"required,number"`
}
