package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
)

type MeasurementController struct {
	measurementService app.MeasurementService
}

func NewMeasurementController(ms app.MeasurementService) MeasurementController {
	return MeasurementController{
		measurementService: ms,
	}
}

func (c MeasurementController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		measurement, err := requests.Bind(r, requests.MeasurementRequest{}, domain.Measurement{})
		if err != nil {
			log.Printf("MeasurementController.Save(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		room := r.Context().Value(RoomKey).(domain.Room)
		measurement.Room_id = room.Id

		measurement, err = c.measurementService.Save(measurement)
		if err != nil {
			log.Printf("MeasurementController.Save(c.measurementService.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		noContent(w)
	}
}
