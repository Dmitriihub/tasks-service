package task

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTask(task Task) (*Task, error) {
	err := s.repo.Create(&task)
	return &task, err
}

func (s *Service) GetTaskByID(id uint) (*Task, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllTasks() ([]Task, error) {
	return s.repo.GetAll()
}

func (s *Service) GetTasksByUser(userID uint) ([]Task, error) {
	return s.repo.GetByUserID(userID)
}

func (s *Service) UpdateTask(task Task) (*Task, error) {
	err := s.repo.Update(&task)
	return &task, err
}

func (s *Service) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
