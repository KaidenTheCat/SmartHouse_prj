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

		rooms, err := c.houseServise.FindRoomByHouseId(house.Id)
		if err != nil {
			log.Printf("HouseController.Find(c.houseServise.FindRoomByHouseId): %s", err)
			return
		}

		var houseFindDto resources.HouseFindDto
		houseFindDto = houseFindDto.DomainToFindDto(house)

		roomsDto := resources.RoomFindListDto{}.DomainToDtoCollection(rooms)
		houseFindDto.Rooms = roomsDto

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

func (c HouseController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		updateHouse, err := requests.Bind(r, requests.UpdateHouseRequest{}, domain.House{})
		if err != nil {
			log.Printf("HouseController.Update(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		houseId := r.Context().Value(HouseKey).(domain.House).Id
		house, err := c.houseServise.FindById(houseId)

		if err != nil {
			log.Printf("HouseController.Update(c.houseServise.FindById): %s", err)
			return
		}

		if updateHouse.Name != "" {
			house.Name = updateHouse.Name
		}
		if updateHouse.City != "" {
			house.City = updateHouse.City
		}
		if updateHouse.Address != "" {
			house.Address = updateHouse.Address
		}
		if updateHouse.Lat != 0 {
			house.Lat = updateHouse.Lat
		}
		if updateHouse.Lon != 0 {
			house.Lon = updateHouse.Lon
		}

		house, err = c.houseServise.Update(house)
		if err != nil {
			log.Printf("HouseController.Update(c.houseServise.Update): %s", err)
			InternalServerError(w, err)
			return
		}

		var houseUpdateDto resources.HouseDto
		houseUpdateDto = houseUpdateDto.DomainToDto(house)
		Success(w, houseUpdateDto)
	}
}

func (c HouseController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		house := r.Context().Value(HouseKey).(domain.House)

		err := c.houseServise.Delete(house.Id)
		if err != nil {
			log.Printf("HouseController.Delete(c.houseServise.Delete):  %s", err)
			InternalServerError(w, err)
			return
		}

		noContent(w)
	}
}
