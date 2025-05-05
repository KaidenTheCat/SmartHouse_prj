package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type RoomDto struct {
	Id          uint64     `json:"id"`
	HouseId     uint64     `json:"houseId"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	CreateDate  time.Time  `json:"createDate"`
	UpdatedDate time.Time  `json:"updateDate"`
	DeleteDate  *time.Time `json:"deleteDate,omitempty"`
}

type RoomFindDto struct {
	Id          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreateDate  time.Time `json:"createDate"`
}

type RoomFindListDto struct {
	Id          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type RoomUpdateDto struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

func (d RoomDto) DomainToDto(r domain.Room) RoomDto {
	return RoomDto{
		Id:          r.Id,
		HouseId:     r.HouseId,
		Name:        r.Name,
		Description: r.Description,
		CreateDate:  r.CreateDate,
		UpdatedDate: r.UpdatedDate,
		DeleteDate:  r.DeleteDate,
	}
}

func (d RoomFindDto) DomainToFindDto(r domain.Room) RoomFindDto {
	return RoomFindDto{
		Id:          r.Id,
		Name:        r.Name,
		Description: r.Description,
		CreateDate:  r.CreateDate,
	}
}

func (d RoomFindListDto) DomainToFindListDto(r domain.Room) RoomFindListDto {
	return RoomFindListDto{
		Id:          r.Id,
		Name:        r.Name,
		Description: r.Description,
	}
}

func (d RoomFindListDto) DomainToDtoCollection(rooms []domain.Room) []RoomFindListDto {
	rm := make([]RoomFindListDto, len(rooms))
	for i, h := range rooms {
		rm[i] = d.DomainToFindListDto(h)
	}

	return rm
}

func (d RoomUpdateDto) DomainToUpdateDto(r domain.Room) RoomUpdateDto {
	return RoomUpdateDto{
		Name:        r.Name,
		Description: r.Description,
	}
}
