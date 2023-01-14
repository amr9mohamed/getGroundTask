package guests

type CreateRequest struct {
	Name         string `json:"name" binding:"required"`
	Table        uint   `json:"table" binding:"required"`
	Accompanying int64  `json:"accompanying_guests" binding:"required" gt:"0"`
}

type CreateResponse struct {
	Name string `json:"name"`
}

type ListDTO struct {
	Guests []GuestListDTO `json:"guests"`
}

type GuestListDTO struct {
	Name         string `json:"name"`
	Table        uint   `json:"table"`
	Accompanying int64  `json:"accompanying_guests"`
}

type DTO struct {
	Guests []GuestDTO `json:"guests"`
}

type GuestDTO struct {
	Name         string `json:"name"`
	Accompanying int64  `json:"accompanying_guests"`
	TimeArrived  string `json:"time_arrived"`
}

type CheckInRequest struct {
	Name         string `json:"name" binding:"required"`
	Accompanying int64  `json:"accompanying_guests" binding:"required" gt:"0"`
}

type CheckInResponse struct {
	Name string `json:"name"`
}
