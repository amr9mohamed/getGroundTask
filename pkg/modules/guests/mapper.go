package guests

import (
	"github.com/getground/tech-tasks/backend/definitions/guests"
)

func mapGuestsListToDTO(gs []guests.Guest) guests.ListDTO {
	list := make([]guests.GuestListDTO, 0, len(gs))

	for _, g := range gs {
		list = append(list, mapGuestListToDTO(g))
	}
	return guests.ListDTO{Guests: list}
}

func mapGuestListToDTO(g guests.Guest) guests.GuestListDTO {
	return guests.GuestListDTO{
		Name:         g.Name,
		Table:        g.TableID,
		Accompanying: g.Accompanying,
	}
}

func mapGuestsToDTO(gs []guests.Guest) guests.DTO {
	list := make([]guests.GuestDTO, 0, len(gs))
	for _, g := range gs {
		list = append(list, mapGuestToDTO(g))
	}
	return guests.DTO{Guests: list}
}

func mapGuestToDTO(g guests.Guest) guests.GuestDTO {
	return guests.GuestDTO{
		Name:         g.Name,
		Accompanying: g.Accompanying,
		TimeArrived:  g.TimeArrived.String(),
	}
}
