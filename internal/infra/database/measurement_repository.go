package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/google/uuid"
	"github.com/upper/db/v4"
)

const MeasurementTableName = "measurements"

type measurement struct {
	Id          uint64    `db:"id,omitempty"`
	Device_uuid uuid.UUID `db:"device_uuid"`
	Room_id     uint64    `db:"room_id"`
	Value       string    `db:"value"`
	CreateDate  time.Time `db:"created_date"`
}

type MeasurementRepository interface {
	Save(m domain.Measurement) (domain.Measurement, error)
}

type measurementRepository struct {
	coll db.Collection
	sess db.Session
}

func NewMeasurementRepository(sess db.Session) MeasurementRepository {
	return measurementRepository{
		coll: sess.Collection(MeasurementTableName),
		sess: sess,
	}
}

func (r measurementRepository) Save(m domain.Measurement) (domain.Measurement, error) {
	ms := r.mapDomainToModel(m)
	ms.CreateDate = time.Now()

	err := r.coll.InsertReturning(&ms)
	if err != nil {
		return domain.Measurement{}, err
	}

	m = r.mapModelToDomain(ms)
	return m, nil
}

func (r measurementRepository) mapDomainToModel(d domain.Measurement) measurement {
	return measurement{
		Id:          d.Id,
		Device_uuid: d.Device_uuid,
		Room_id:     d.Room_id,
		Value:       d.Value,
		CreateDate:  d.CreateDate,
	}
}

func (r measurementRepository) mapModelToDomain(m measurement) domain.Measurement {
	return domain.Measurement{
		Id:          m.Id,
		Device_uuid: m.Device_uuid,
		Room_id:     m.Room_id,
		Value:       m.Value,
		CreateDate:  m.CreateDate,
	}
}

func (r measurementRepository) mapModelToDomainCollection(ms []measurement) []domain.Measurement {
	dm := make([]domain.Measurement, len(ms))
	for i, h := range ms {
		dm[i] = r.mapModelToDomain(h)
	}
	return dm
}
