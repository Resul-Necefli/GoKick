package repository

import (
	"errors"

	"github.com/Resul-Necefli/GoKick/internal/domain"
)

var (
	ErrNotFound  = errors.New("campaign not found")
	ErrDuplicate = errors.New("campaign with given ID already exists")
)

type InMemoryCampaignRepository struct {
	data map[int]*domain.Campaign
}

func (i *InMemoryCampaignRepository) Create(c *domain.Campaign) error {

	if _, k := i.data[c.ID]; k {
		return ErrDuplicate
	}

	i.data[c.ID] = c
	return nil
}

func (i *InMemoryCampaignRepository) FindByID(id int) (*domain.Campaign, error) {

	if v := i.data[id]; v != nil {
		return v, nil
	}

	return nil, ErrNotFound

}

func (i *InMemoryCampaignRepository) FindAll() map[int]*domain.Campaign {

	return i.data
}

func (i *InMemoryCampaignRepository) Update(c *domain.Campaign) (*domain.Campaign, error) {

	if v := i.data[c.ID]; v != nil {
		i.data[c.ID] = c

		return i.data[c.ID], nil
	}

	return nil, ErrNotFound

}
