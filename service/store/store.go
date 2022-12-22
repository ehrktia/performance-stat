//go:generate mockgen --destination=./mocks/${GOFILE} --package=mocks --source=${GOFILE}
package store

type Repository interface {
	GetByID(id int) ([]byte, error)
}

type storeService struct {
	store Repository
}

func (s *storeService) RetreiveNameByID(id int) (string, error) {
	defaultUser := "defaultUser"
	dataBytes, err := s.store.GetByID(id)
	if err != nil {
		return defaultUser, err
	}
	return string(dataBytes), nil

}
