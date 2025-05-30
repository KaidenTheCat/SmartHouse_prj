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

		if (device.Power_consumption != nil && device.Units != nil) || // Не може бути і те і те
			(device.Category != domain.SENSOR && device.Category != domain.ACTUATOR) || //Повинно бути або SENSOR або ACTUATOR
			(device.Category == domain.SENSOR && device.Units == nil) || //якщо є SENSOR то має бути Units
			(device.Category == domain.ACTUATOR && device.Power_consumption == nil) { //якщо є ACTUATOR то має бути Power_consumption

			err = errors.New("wrong data request")
			Forbidden(w, err)
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

func (c DeviceController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		updateDevice, err := requests.Bind(r, requests.UpdateDeviceRequest{}, domain.Device{})
		if err != nil {
			log.Printf("DeviceController.Update(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		deviceId := r.Context().Value(DeviceKey).(domain.Device).Id
		device, err := c.deviceServise.FindById(deviceId)

		if err != nil {
			log.Printf("DeviceController.Update(c.deviceServise.FindById): %s", err)
			return
		}

		if updateDevice.Serial_number != "" {
			device.Serial_number = updateDevice.Serial_number
		}
		if updateDevice.Category != "" {
			device.Category = updateDevice.Category
		}
		if updateDevice.Characteristics != nil {
			device.Characteristics = updateDevice.Characteristics
		}

		if updateDevice.House_id != 0 {
			device.House_id = updateDevice.House_id
		}
		if updateDevice.Room_id != 0 {
			device.Room_id = updateDevice.Room_id
		}

		if device.Category == domain.SENSOR {
			if updateDevice.Units != nil {
				device.Units = updateDevice.Units
			}
			device.Power_consumption = nil
		}

		if device.Category == domain.ACTUATOR {
			if updateDevice.Power_consumption != nil {
				device.Power_consumption = updateDevice.Power_consumption
			}
			device.Units = nil
		}

		device, err = c.deviceServise.Update(device)
		if err != nil {
			log.Printf("DeviceController.Update(c.deviceServise.Update): %s", err)
			InternalServerError(w, err)
			return
		}

		var deviceUpdateDto resources.DeviceDto
		deviceUpdateDto = deviceUpdateDto.DomainToDto(device)
		Success(w, deviceUpdateDto)
	}
}

func (c DeviceController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		device := r.Context().Value(DeviceKey).(domain.Device)

		err := c.deviceServise.Delete(device.Id)
		if err != nil {
			log.Printf("DeviceController.Delete(c.deviceServise.Delete):  %s", err)
			InternalServerError(w, err)
			return
		}

		noContent(w)
	}
}

func (c DeviceController) Move() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		house_id := r.Context().Value(HouseKey).(domain.House).Id
		room_id := r.Context().Value(RoomKey).(domain.Room).Id
		device := r.Context().Value(DeviceKey).(domain.Device)

		device.House_id = house_id
		device.Room_id = room_id

		device, err := c.deviceServise.Update(device)
		if err != nil {
			log.Printf("DeviceController.Move(c.deviceServise.Update):  %s", err)
			InternalServerError(w, err)
			return
		}

		var deviceUpdateDto resources.DeviceDto
		deviceUpdateDto = deviceUpdateDto.DomainToDto(device)
		Success(w, deviceUpdateDto)
	}
}
