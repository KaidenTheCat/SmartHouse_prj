package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const RoomsTableName = "rooms"

type room struct {
	Id          uint64     `db:"id,omitempty"`
	HouseId     uint64     `db:"house_id"`
	Name        string     `db:"name"`
	Description *string    `db:"description"`
	CreateDate  time.Time  `db:"created_date"`
	UpdatedDate time.Time  `db:"updated_date"`
	DeleteDate  *time.Time `db:"deleted_date"`
}

type RoomRepository interface {
	Save(dr domain.Room) (domain.Room, error)
	Find(id uint64) (domain.Room, error)
	FindList(hId uint64) ([]domain.Room, error)
	Update(rm domain.Room) (domain.Room, error)
	Delete(id uint64) error
}

type roomRepository struct {
	coll db.Collection
	sess db.Session
}

func NewRoomRepository(sess db.Session) RoomRepository {
	return roomRepository{
		coll: sess.Collection(RoomsTableName),
		sess: sess,
	}
}

func (r roomRepository) Save(dr domain.Room) (domain.Room, error) {
	rm := r.mapDomainToModel(dr)
	rm.CreateDate = time.Now()
	rm.UpdatedDate = time.Now()

	err := r.coll.InsertReturning(&rm)
	if err != nil {
		return domain.Room{}, err
	}

	dr = r.mapModelToDomain(rm)
	return dr, nil
}

func (r roomRepository) Find(id uint64) (domain.Room, error) {
	var rm room
	err := r.coll.
		Find(db.Cond{
			"id":           id,
			"deleted_date": nil}).One(&rm)
	if err != nil {
		return domain.Room{}, err
	}

	hs := r.mapModelToDomain(rm)
	return hs, nil
}

func (r roomRepository) FindList(hId uint64) ([]domain.Room, error) {
	var rooms []room
	err := r.coll.
		Find(db.Cond{
			"house_id":     hId,
			"deleted_date": nil}).All(&rooms)
	if err != nil {
		return nil, err
	}

	rm := r.mapModelToDomainCollection(rooms)
	return rm, nil
}

func (r roomRepository) Update(rm domain.Room) (domain.Room, error) {
	room := r.mapDomainToModel(rm)
	room.UpdatedDate = time.Now()

	err := r.coll.UpdateReturning(&room)
	if err != nil {
		return domain.Room{}, err
	}

	rm = r.mapModelToDomain(room)
	return rm, nil
}

func (r roomRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r roomRepository) mapModelToDomain(rm room) domain.Room {
	return domain.Room{
		Id:          rm.Id,
		HouseId:     rm.HouseId,
		Name:        rm.Name,
		Description: rm.Description,
		CreateDate:  rm.CreateDate,
		UpdatedDate: rm.UpdatedDate,
		DeleteDate:  rm.DeleteDate,
	}
}

func (r roomRepository) mapDomainToModel(d domain.Room) room {
	return room{
		Id:          d.Id,
		HouseId:     d.HouseId,
		Name:        d.Name,
		Description: d.Description,
		CreateDate:  d.CreateDate,
		UpdatedDate: d.UpdatedDate,
		DeleteDate:  d.DeleteDate,
	}
}

func (r roomRepository) mapModelToDomainCollection(rooms []room) []domain.Room {
	rm := make([]domain.Room, len(rooms))
	for i, h := range rooms {
		rm[i] = r.mapModelToDomain(h)
	}
	return rm
}
