package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

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
		// var houseSaveDto resources.HouseSaveDto
		// houseSaveDto = houseSaveDto.DomainToSaveDto(house)
		// Success(w, resources.UserDto{}.DomainToDto(user))

		var houseSaveDto resources.HouseSaveDto
		houseSaveDto = houseSaveDto.DomainToSaveDto(house)
		Success(w, houseSaveDto)
	}
}

func (c HouseController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// house, err := requests.Bind(r, requests.HouseRequest{}, domain.House{})
		// if err != nil {
		// 	log.Printf("HouseController.Find(requests.Bind): %s", err)
		// 	BadRequest(w, errors.New("invalid request body"))
		// 	return
		// }
		idStr := r.URL.Query().Get("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			log.Printf("HouseController.Find(strconv.ParseUint): %s", err)
			BadRequest(w, errors.New("invalid id"))
			return
		}

		house, err := c.houseServise.Find(id)
		if err != nil {
			log.Printf("HouseController.Find(c.houseServise.Find): %s", err)
			InternalServerError(w, err)
			return
		}
		var houseFindDto resources.HouseFindDto
		houseFindDto = houseFindDto.DomainToFindDto(house)
		Success(w, houseFindDto)
	}
}

func (c HouseController) FindList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		uId := r.Context().Value(UserKey).(domain.User).Id

		houses, err := c.houseServise.FindList(uId)
		if err != nil {
			log.Printf("HouseController.FindList(c.houseServise.FindList): %s", err)
			InternalServerError(w, err)
			return
		}
		housesFindListDto := c.mapDomainToFindListDto(houses)
		Success(w, housesFindListDto)
	}
}

func (c HouseController) mapDomainToFindListDto(houses []domain.House) []resources.HouseFindDto {
	hs := make([]resources.HouseFindDto, len(houses))
	for i, house := range houses {
		hs[i] = resources.HouseFindDto{}.DomainToFindDto(house)
	}
	return hs
}
