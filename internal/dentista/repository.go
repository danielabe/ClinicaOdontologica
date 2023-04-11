package dentista

import (
	"clinicaodontologica/internal/domain"
	"clinicaodontologica/pkg/store"
	"errors"
)

type IRepository interface {
	GetById(id int) (domain.Dentista, error)
	Create(d domain.Dentista) (domain.Dentista, error)
	Update(d domain.Dentista, id int) (domain.Dentista, error)
	Delete(id int) error
}

type Repository struct {
	Store store.StoreInterface
}

func NewRepository(Store store.StoreInterface) IRepository {
	return &Repository{Store}
}

func (r *Repository) GetById(id int) (domain.Dentista, error) {
	dent, err := r.Store.ReadDentista(id)
	if err != nil {
		return domain.Dentista{}, errors.New("Dentista not found")
	}

	return dent, nil
}

func (r *Repository) Create(d domain.Dentista) (domain.Dentista, error) {
	if r.Store.ExistsMatricula(d.Matricula) {
		return domain.Dentista{}, errors.New("Matricula value already exists")
	}
	dent, err := r.Store.CreateDentista(d)
	if err != nil {
		return domain.Dentista{}, errors.New("Error creating dentista")
	}
	return dent, nil
}

func (r *Repository) Update(d domain.Dentista, id int) (domain.Dentista, error) {
	if r.Store.ExistsMatriculaWithOtherId(d.Matricula, id) {
		return domain.Dentista{}, errors.New("Matricula value already exists")
	}
	dent, err := r.Store.UpdateDentista(d, id)
	if err != nil {
		return domain.Dentista{}, errors.New("Error updating dentista")
	}
	return dent, nil
}

func (r *Repository) Delete(id int) error {
	if !r.Store.ExistsDentistaId(id) {
		return errors.New("Dentista id does not exist")
	}
	err := r.Store.DeleteDentista(id)
	if err != nil {
		return err
	}

	return nil
}
