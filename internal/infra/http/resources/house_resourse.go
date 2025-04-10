package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type HouseSaveDto struct {
	Id          uint64     `json:"id"`
	UserId      uint64     `json:"userId"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	City        string     `json:"city"`
	Address     string     `json:"address"`
	Lat         float64    `json:"lat"`
	Lon         float64    `json:"lon"`
	CreateDate  time.Time  `json:"createDate"`
	UpdatedDate time.Time  `json:"updateDate"`
	DeleteDate  *time.Time `json:"deleteDate,omitempty"`
}

type HouseFindDto struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	City        string  `json:"city"`
	Address     string  `json:"address"`
}

func (d HouseSaveDto) DomainToSaveDto(h domain.House) HouseSaveDto {
	return HouseSaveDto{
		Id:          h.Id,
		UserId:      h.UserId,
		Name:        h.Name,
		Description: h.Description,
		City:        h.City,
		Address:     h.Address,
		Lat:         h.Lat,
		Lon:         h.Lon,
		CreateDate:  h.CreateDate,
		UpdatedDate: h.UpdatedDate,
		DeleteDate:  h.DeleteDate,
	}
}

func (d HouseFindDto) DomainToFindDto(h domain.House) HouseFindDto {
	return HouseFindDto{
		Name:        h.Name,
		Description: h.Description,
		City:        h.City,
		Address:     h.Address,
	}
}
