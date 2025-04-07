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
		var houseDto resources.HouseDto
		houseDto = houseDto.DomainToDto(house)
		Success(w, resources.UserDto{}.DomainToDto(user))
	}
}
