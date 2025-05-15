package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/google/uuid"
	"github.com/upper/db/v4"
)

const EventTableName = "measurements"

type event struct {
	Id          uint64    `db:"id,omitempty"`
	Device_uuid uuid.UUID `db:"device_uuid"`
	Room_id     uint64    `db:"room_id"`
	Action      string    `db:"value"`
	CreateDate  time.Time `db:"created_date"`
}

type EventRepository interface {
	Save(e domain.Event) (domain.Event, error)
}

type eventRepository struct {
	coll db.Collection
	sess db.Session
}

func NewEventRepository(sess db.Session) EventRepository {
	return eventRepository{
		coll: sess.Collection(EventTableName),
		sess: sess,
	}
}

func (r eventRepository) Save(e domain.Event) (domain.Event, error) {
	ev := r.mapDomainToModel(e)
	ev.CreateDate = time.Now()

	err := r.coll.InsertReturning(&ev)
	if err != nil {
		return domain.Event{}, err
	}

	e = r.mapModelToDomain(ev)
	return e, nil
}

func (r eventRepository) mapDomainToModel(d domain.Event) event {
	return event{
		Id:          d.Id,
		Device_uuid: d.Device_uuid,
		Action:      d.Action,
		Room_id:     d.Room_id,
		CreateDate:  d.CreateDate,
	}
}

func (r eventRepository) mapModelToDomain(e event) domain.Event {
	return domain.Event{
		Id:          e.Id,
		Device_uuid: e.Device_uuid,
		Room_id:     e.Room_id,
		Action:      e.Action,
		CreateDate:  e.CreateDate,
	}
}

func (r eventRepository) mapModelToDomainCollection(ev []event) []domain.Event {
	de := make([]domain.Event, len(ev))
	for i, h := range ev {
		de[i] = r.mapModelToDomain(h)
	}
	return de
}
