package task

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTask(t Task) (*Task, error) {
	err := s.repo.Create(&t)
	return &t, err
}

func (s *Service) GetTaskByID(id uint) (*Task, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllTasks() ([]Task, error) {
	return s.repo.GetAll()
}

func (s *Service) GetTasksByUser(userID uint) ([]Task, error) {
	return s.repo.GetByUser(userID)
}

func (s *Service) UpdateTask(t *Task) (*Task, error) {
	err := s.repo.Update(t)
	return t, err
}

func (s *Service) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
