package repository

import "github.com/Resul-Necefli/GoKick/internal/domain"

type CampaignRepository interface {
	Create(c *domain.Campaign) error
	FindByID(id int) (*domain.Campaign, error)
	FindAll() (map[int]*domain.Campaign, error)
	Update(c *domain.Campaign) (*domain.Campaign, error)
}
