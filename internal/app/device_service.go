package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type DeviceService interface {
	Save(h domain.Device) (domain.Device, error)
	Find(id uint64) (interface{}, error)
	FindById(id uint64) (domain.Device, error)
	Delete(id uint64) error
}

type deviceService struct {
	deviceRepo database.DeviceRepository
}

func NewHDeviceService(dr database.DeviceRepository) deviceService {
	return deviceService{
		deviceRepo: dr,
	}
}

func (s deviceService) Save(h domain.Device) (domain.Device, error) {
	device, err := s.deviceRepo.Save(h)
	if err != nil {
		log.Printf("deviceService.Save(s.deviceRepo.Save): %s", err)
		return domain.Device{}, err
	}

	return device, nil
}

func (s deviceService) Find(id uint64) (interface{}, error) {
	device, err := s.deviceRepo.Find(id)
	if err != nil {
		log.Printf("deviceService.Find(s.deviceRepo.Find): %s", err)
		return domain.Device{}, err
	}

	return device, nil
}

func (s deviceService) FindById(id uint64) (domain.Device, error) {
	device, err := s.deviceRepo.Find(id)
	if err != nil {
		log.Printf("deviceService.FindById(s.deviceRepo.Find): %s", err)
		return domain.Device{}, err
	}

	return device, nil
}

func (s deviceService) Delete(id uint64) error {
	err := s.deviceRepo.Delete(id)
	if err != nil {
		log.Printf("deviceService.Delete(s.deviceRepo.Delete): %s", err)
		return err
	}

	return nil
}
