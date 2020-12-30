package web

type ModelDetailRequest struct {
	ID int `uri:"id" binding:"required"`
}
