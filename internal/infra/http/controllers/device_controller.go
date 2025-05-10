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

type DeviceController struct {
	deviceServise app.DeviceService
}

func NewDeviceController(ds app.DeviceService) DeviceController {
	return DeviceController{
		deviceServise: ds,
	}
}

func (c DeviceController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		device, err := requests.Bind(r, requests.DeviceRequest{}, domain.Device{})
		if err != nil {
			log.Printf("DeviceController.Save(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		house := r.Context().Value(HouseKey).(domain.House)
		room := r.Context().Value(RoomKey).(domain.Room)

		device.House_id = house.Id
		device.Room_id = room.Id

		device, err = c.deviceServise.Save(device)
		if err != nil {
			log.Printf("DeviceController.Save(c.deviceServise.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		var deviceSaveDto resources.DeviceDto
		deviceSaveDto = deviceSaveDto.DomainToDto(device)
		Success(w, deviceSaveDto)
	}
}

func (c DeviceController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		device := r.Context().Value(DeviceKey).(domain.Device)

		var deviceFindDto resources.DeviceFindDto
		deviceFindDto = deviceFindDto.DomainToFindDto(device)
		Success(w, deviceFindDto)
	}
}
