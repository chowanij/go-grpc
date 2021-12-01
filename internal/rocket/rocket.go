package rocket

import "context"

// Rocket - contains fileld definition
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

// Store - defines the database access implementation
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service - our service - responsible for updating rockets
type Service struct {
	Store Store
}

// New - returns instance of new service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// GetRocketByID - retrieves a rocket based on the ID from the store
func (s Service)GetRocketByID(ctx context.Context, id string)(Rocket, error){
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

// InsertRocket - insert new rocket into store
func (s Service)InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

func (s Service)DeleteRocket(id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}


