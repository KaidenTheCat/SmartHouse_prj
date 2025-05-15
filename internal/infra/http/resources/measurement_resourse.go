// package resources

// import (
// 	"time"

// 	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
// )

// type MeasurementDto struct {
// 	Value      string    `json:"value"`
// 	Room_id    uint64    `json:"room_id"`
// 	CreateDate time.Time `json:"createDate"`
// }

// func (d MeasurementDto) DomainToDto(dm domain.Measurement) MeasurementDto {
// 	return MeasurementDto{
// 		Value:      dm.Value,
// 		Room_id:    dm.Room_id,
// 		CreateDate: dm.CreateDate,
// 	}
// }
