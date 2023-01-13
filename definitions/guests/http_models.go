package guests

type CreateRequest struct {
	Name         string `json:"name" binding:"required"`
	Table        uint   `json:"table" binding:"required"`
	Accompanying int64  `json:"accompanying_guests" binding:"required"`
}

type CreateResponse struct {
	Name string `json:"name"`
}
