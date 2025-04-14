package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type HouseController struct {
	houseServise app.HouseService
}

func NewHouseController(hs app.HouseService) HouseController {
	return HouseController{
		houseServise: hs,
	}
}

func (c HouseController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		house, err := requests.Bind(r, requests.HouseRequest{}, domain.House{})
		if err != nil {
			log.Printf("HouseController.Save(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		house.UserId = user.Id

		house, err = c.houseServise.Save(house)
		if err != nil {
			log.Printf("HouseController.Save(c.houseServise.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		var houseSaveDto resources.HouseDto
		houseSaveDto = houseSaveDto.DomainToDto(house)
		Success(w, houseSaveDto)
	}
}

func (c HouseController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		house := r.Context().Value(HouseKey).(domain.House)
		user := r.Context().Value(UserKey).(domain.User)

		if house.UserId != user.Id {
			err := errors.New("access denied")
			Forbidden(w, err)
			return
		}

		var houseFindDto resources.HouseFindDto
		houseFindDto = houseFindDto.DomainToFindDto(house)
		Success(w, houseFindDto)
	}
}

func (c HouseController) FindList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		houses, err := c.houseServise.FindList(user.Id)

		if err != nil {
			log.Printf("HouseController.FindList(c.houseServise.FindList): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.HouseFindListDto{}.DomainToDtoCollection(houses))
	}
}
