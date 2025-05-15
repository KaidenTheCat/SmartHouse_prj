package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type MeasurementService interface {
	Save(m domain.Measurement) (domain.Measurement, error)
}

type measurementService struct {
	measurementRepo database.MeasurementRepository
}

func NewMeasurementService(mr database.MeasurementRepository) measurementService {
	return measurementService{
		measurementRepo: mr,
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
