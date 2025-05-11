package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/google/uuid"
	"github.com/upper/db/v4"
)

const DeviceTableName = "devices"

type device struct {
	Id                uint64     `db:"id,omitempty"`
	House_id          uint64     `db:"house_id"`
	Room_id           uint64     `db:"room_id"`
	Uuid              uuid.UUID  `db:"uuid"`
	Serial_number     string     `db:"serial_number"`
	Characteristics   *string    `db:"characteristics"`
	Category          string     `db:"category"`
	Units             *string    `db:"units"`
	Power_consumption *string    `db:"power_consumption"`
	CreateDate        time.Time  `db:"created_date"`
	UpdatedDate       time.Time  `db:"updated_date"`
	DeleteDate        *time.Time `db:"deleted_date"`
}

type DeviceRepository interface {
	Save(d domain.Device) (domain.Device, error)
	Find(id uint64) (domain.Device, error)
	FindList(dId uint64) ([]domain.Device, error)
	Update(d domain.Device) (domain.Device, error)
	Delete(id uint64) error
}

type deviceRepository struct {
	coll db.Collection
	sess db.Session
}

func NewDeviceRepository(sess db.Session) DeviceRepository {
	return deviceRepository{
		coll: sess.Collection(DeviceTableName),
		sess: sess,
	}
}

func (r deviceRepository) Save(d domain.Device) (domain.Device, error) {
	dv := r.mapDomainToModel(d)
	dv.CreateDate = time.Now()
	dv.UpdatedDate = time.Now()
	dv.Uuid = uuid.New()

	err := r.coll.InsertReturning(&dv)
	if err != nil {
		return domain.Device{}, err
	}

	d = r.mapModelToDomain(dv)
	return d, nil
}

func (r deviceRepository) Find(id uint64) (domain.Device, error) {
	var d device
	err := r.coll.
		Find(db.Cond{
			"id":           id,
			"deleted_date": nil}).One(&d)
	if err != nil {
		return domain.Device{}, err
	}

	dv := r.mapModelToDomain(d)
	return dv, nil
}

func (r deviceRepository) FindList(rId uint64) ([]domain.Device, error) {
	var d []device
	err := r.coll.
		Find(db.Cond{
			"room_id":      rId,
			"deleted_date": nil}).All(&d)
	if err != nil {
		return nil, err
	}

	dv := r.mapModelToDomainCollection(d)
	return dv, nil
}
func (r deviceRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r deviceRepository) Update(d domain.Device) (domain.Device, error) {
	dv := r.mapDomainToModel(d)
	dv.UpdatedDate = time.Now()

	err := r.coll.UpdateReturning(&dv)
	if err != nil {
		return domain.Device{}, err
	}

	d = r.mapModelToDomain(dv)
	return d, nil
}

func (r deviceRepository) mapDomainToModel(d domain.Device) device {
	return device{
		Id:                d.Id,
		House_id:          d.House_id,
		Room_id:           d.Room_id,
		Uuid:              d.Uuid,
		Serial_number:     d.Serial_number,
		Characteristics:   d.Characteristics,
		Category:          d.Category,
		Units:             d.Units,
		Power_consumption: d.Power_consumption,
		CreateDate:        d.CreateDate,
		UpdatedDate:       d.UpdatedDate,
		DeleteDate:        d.DeleteDate,
	}
}

func (r deviceRepository) mapModelToDomain(d device) domain.Device {
	return domain.Device{
		Id:                d.Id,
		House_id:          d.House_id,
		Room_id:           d.Room_id,
		Uuid:              d.Uuid,
		Serial_number:     d.Serial_number,
		Characteristics:   d.Characteristics,
		Category:          d.Category,
		Units:             d.Units,
		Power_consumption: d.Power_consumption,
		CreateDate:        d.CreateDate,
		UpdatedDate:       d.UpdatedDate,
		DeleteDate:        d.DeleteDate,
	}
}

func (r deviceRepository) mapModelToDomainCollection(devices []device) []domain.Device {
	dm := make([]domain.Device, len(devices))
	for i, h := range devices {
		dm[i] = r.mapModelToDomain(h)
	}
	return dm
}
