package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type HouseDto struct {
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
	Id          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	City        string    `json:"city"`
	Address     string    `json:"address"`
	CreateDate  time.Time `json:"createDate"`
}

type HouseFindListDto struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	City        string  `json:"city"`
	Address     string  `json:"address"`
}

func (d HouseDto) DomainToDto(h domain.House) HouseDto {
	return HouseDto{
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
		Id:          h.Id,
		Name:        h.Name,
		Description: h.Description,
		City:        h.City,
		Address:     h.Address,
		CreateDate:  h.CreateDate,
	}
}

func (d HouseFindListDto) DomainToFindListDto(h domain.House) HouseFindListDto {
	return HouseFindListDto{
		Name:        h.Name,
		Description: h.Description,
		City:        h.City,
		Address:     h.Address,
	}
}

func (d HouseFindListDto) DomainToDtoCollection(houses []domain.House) []HouseFindListDto {
	hs := make([]HouseFindListDto, len(houses))
	for i, h := range houses {
		hs[i] = d.DomainToFindListDto(h)
	}

	return hs
}
