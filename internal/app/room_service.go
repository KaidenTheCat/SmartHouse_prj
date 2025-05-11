package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type RoomService interface {
	Save(h domain.Room) (domain.Room, error)
	Find(id uint64) (interface{}, error)
	FindById(id uint64) (domain.Room, error)
	FindDeviceByRoomId(id uint64) ([]domain.Device, error)
	Update(r domain.Room) (domain.Room, error)
	Delete(id uint64) error
}

type roomService struct {
	roomRepo   database.RoomRepository
	deviceRepo database.DeviceRepository
}

func NewRoomService(hr database.RoomRepository, dr database.DeviceRepository) roomService {
	return roomService{
		roomRepo:   hr,
		deviceRepo: dr,
	}
}

func (s roomService) Save(h domain.Room) (domain.Room, error) {
	room, err := s.roomRepo.Save(h)
	if err != nil {
		log.Printf("roomService.Save(s.roomRepo.Save): %s", err)
		return domain.Room{}, err
	}

	return room, nil
}

func (s roomService) Find(id uint64) (interface{}, error) {
	room, err := s.roomRepo.Find(id)
	if err != nil {
		log.Printf("roomService.Find(s.roomRepo.Find): %s", err)
		return domain.Room{}, err
	}

	return room, nil
}

func (s roomService) FindDeviceByRoomId(id uint64) ([]domain.Device, error) {
	device, err := s.deviceRepo.FindList(id)
	if err != nil {
		log.Printf("roomService.FindDeviceByRoomId(s.deviceRepo.FindList): %s", err)
		return []domain.Device{}, err
	}

	return device, nil
}

func (s roomService) FindById(id uint64) (domain.Room, error) {
	room, err := s.roomRepo.Find(id)
	if err != nil {
		log.Printf("roomService.FindById(s.roomRepo.Find): %s", err)
		return domain.Room{}, err
	}

	return room, nil
}

func (s roomService) Update(r domain.Room) (domain.Room, error) {
	room, err := s.roomRepo.Update(r)
	if err != nil {
		log.Printf("roomService.Update(.roomRepo.Update): %s", err)
		return domain.Room{}, err
	}

	return room, nil
}

func (s roomService) Delete(id uint64) error {
	err := s.roomRepo.Delete(id)
	if err != nil {
		log.Printf("roomService.Delete(s.roomRepo.Delete): %s", err)
		return err
	}

	return nil
}
