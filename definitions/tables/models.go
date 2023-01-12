package tables

type Table struct {
	ID         uint `gorm:"primarykey"`
	Capacity   int64
	EmptySeats int64
}
