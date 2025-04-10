package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type HouseService interface {
	Save(h domain.House) (domain.House, error)
	Find(id uint64) (domain.House, error)
	FindList(uId uint64) ([]domain.House, error)
}

type houseService struct {
	houseRepo database.HouseRepository
}

func NewHouseService(hr database.HouseRepository) houseService {
	return houseService{
		houseRepo: hr,
	}
}

func (s houseService) Save(h domain.House) (domain.House, error) {
	house, err := s.houseRepo.Save(h)
	if err != nil {
		log.Printf("houseService.Save(s houseService):%s", err)
		return domain.House{}, err
	}

	return house, nil
}

func (s houseService) Find(id uint64) (domain.House, error) {
	house, err := s.houseRepo.Find(id)
	if err != nil {
		log.Printf("houseService.Find(s houseService):%s", err)
		return domain.House{}, err
	}

	return house, nil
}

func (s houseService) FindList(uId uint64) ([]domain.House, error) {
	house, err := s.houseRepo.FindList(uId)
	if err != nil {
		log.Printf("houseService.FindList(s houseService):%s", err)
		return nil, err
	}

	return house, nil
}
