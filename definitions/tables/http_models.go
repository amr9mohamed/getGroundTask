package tables

type CreateRequest struct {
	Capacity int64 `json:"capacity" binding:"required"`
}

type CreateResponse struct {
	ID       uint  `json:"id"`
	Capacity int64 `json:"capacity"`
}
