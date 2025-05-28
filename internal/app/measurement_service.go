package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"github.com/google/uuid"
)

type MeasurementService interface {
	Save(m domain.Measurement) (domain.Measurement, error)
	FindByUuid(uuid uuid.UUID) (domain.Device, error)
}

type measurementService struct {
	measurementRepo database.MeasurementRepository
	deviceRepo      database.DeviceRepository
}

func NewMeasurementService(mr database.MeasurementRepository, dr database.DeviceRepository) measurementService {
	return measurementService{
		measurementRepo: mr,
		deviceRepo:      dr,
	}
}

func (s measurementService) Save(m domain.Measurement) (domain.Measurement, error) {
	measurement, err := s.measurementRepo.Save(m)
	if err != nil {
		log.Printf("measurementService.Save(s.measurementRepo.Save): %s", err)
		return domain.Measurement{}, err
	}

	return measurement, nil
}

func (s measurementService) FindByUuid(uuid uuid.UUID) (domain.Device, error) {
	device, err := s.deviceRepo.FindByUuid(uuid)
	if err != nil {
		log.Printf("measurementService.FindByUuid(s.deviceRepo.FindByUuid): %s", err)
		return domain.Device{}, err
	}

	return device, nil
}
