package task

type Task struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Title  string
	IsDone bool
}
