package task

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(task *Task) error {
	return r.db.Create(task).Error
}

func (r *Repository) GetByID(id uint) (*Task, error) {
	var task Task
	err := r.db.First(&task, id).Error
	return &task, err
}

func (r *Repository) GetAll() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *Repository) GetByUserID(userID uint) ([]Task, error) {
	var tasks []Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *Repository) Update(task *Task) error {
	return r.db.Save(task).Error
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&Task{}, id).Error
}
