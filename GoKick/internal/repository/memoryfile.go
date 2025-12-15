package repository

import (
	"github.com/Resul-Necefli/GoKick/internal/domain"
)

type InMemoryCampaignRepository struct {
	Data map[int]*domain.Campaign
}

func (i *InMemoryCampaignRepository) Create(c *domain.Campaign) error {

	if _, k := i.Data[c.ID]; k {
		return domain.ErrDuplicate
	}

	i.Data[c.ID] = c
	return nil
}

func (i *InMemoryCampaignRepository) FindByID(id int) (*domain.Campaign, error) {

	if v, ok := i.Data[id]; ok {
		return v, nil
	}

	return nil, domain.ErrNotFound

}

func (i *InMemoryCampaignRepository) FindAll() (map[int]*domain.Campaign, error) {

	if i.Data == nil {
		return nil, domain.ErrNotFound
	}
	return i.Data, nil

}

func (i *InMemoryCampaignRepository) Update(c *domain.Campaign) (*domain.Campaign, error) {

	if _, ok := i.Data[c.ID]; ok {
		i.Data[c.ID] = c

		return i.Data[c.ID], nil
	}

	return nil, domain.ErrNotFound

}
