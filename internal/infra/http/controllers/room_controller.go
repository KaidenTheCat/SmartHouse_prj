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

type RoomController struct {
	roomService app.RoomService
}

func NewRoomController(hs app.RoomService) RoomController {
	return RoomController{
		roomService: hs,
	}
}

func (c RoomController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		room, err := requests.Bind(r, requests.RoomRequest{}, domain.Room{})
		if err != nil {
			log.Printf("RoomController.Save(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		house := r.Context().Value(HouseKey).(domain.House)
		room.HouseId = house.Id

		room, err = c.roomService.Save(room)
		if err != nil {
			log.Printf("RoomController.Save(c.roomServise.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		var roomSaveDto resources.RoomDto
		roomSaveDto = roomSaveDto.DomainToDto(room)
		Success(w, roomSaveDto)
	}
}

func (c RoomController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		room := r.Context().Value(RoomKey).(domain.Room)
		devices, err := c.roomService.FindDeviceByRoomId(room.Id)
		if err != nil {
			log.Printf("RoomController.Find(c.roomService.FindDeviceByRoomId): %s", err)
			return
		}

		var roomFindDto resources.RoomFindDto
		roomFindDto = roomFindDto.DomainToFindDto(room)

		devicesDto := resources.DeviceFindListDto{}.DomainToDtoCollection(devices)
		roomFindDto.Devices = devicesDto

		Success(w, roomFindDto)
	}
}

func (c RoomController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		updateRoom, err := requests.Bind(r, requests.RoomRequest{}, domain.Room{})
		if err != nil {
			log.Printf("RoomController.Update(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		roomId := r.Context().Value(RoomKey).(domain.Room).Id
		room, err := c.roomService.FindById(roomId)
		if err != nil {
			log.Printf("RoomController.Update(c.roomServise.FindById): %s", err)
			return
		}

		room.Name = updateRoom.Name
		if updateRoom.Description != nil {
			room.Description = updateRoom.Description
		}

		room, err = c.roomService.Update(room)
		if err != nil {
			log.Printf("RoomController.Update(c.roomServise.Update): %s", err)
			InternalServerError(w, err)
			return
		}

		var roomUpdateDto resources.RoomUpdateDto
		roomUpdateDto = roomUpdateDto.DomainToUpdateDto(room)
		Success(w, roomUpdateDto)
	}
}

func (c RoomController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		room := r.Context().Value(RoomKey).(domain.Room)

		err := c.roomService.Delete(room.Id)
		if err != nil {
			log.Printf("RoomController.Delete(c.roomServise.Delete):  %s", err)
			InternalServerError(w, err)
			return
		}

		noContent(w)
	}
}
