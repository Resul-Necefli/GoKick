package service

import (
	"errors"

	"github.com/Resul-Necefli/GoKick/internal/domain"
	"github.com/Resul-Necefli/GoKick/internal/repository"
)

type CampaignService struct {
	repo repository.CampaignRepository
}

func NewCampaignService(repo repository.CampaignRepository) *CampaignService {
	return &CampaignService{
		repo: repo,
	}
}

func (c *CampaignService) CreateCampaign(d *domain.Campaign) error {

	if d.ID < 0 {
		return errors.New("ID value cannot be negative")
	}
	if len([]byte(d.Description)) >= 200 {
		return errors.New("description exceeds maximum length of 200 characters")
	}

	if d.Status != domain.CampaignStatusActive {
		return errors.New("campaign status must be active")
	}
	if d.TargetAmount <= 200 {
		return errors.New("target amount must be greater than 200")
	}

	if d.CurrentAmount < 0 {
		return errors.New("current amount cannot be negative")
	}

	err := c.repo.Create(d)
	return err

}

func (c *CampaignService) GetCampaignByID(id int) (*domain.Campaign, error) {

	if id < 0 {
		return nil, errors.New("ID value cannot be negative")
	}

	campaign, err := c.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return campaign, nil

}
func (c *CampaignService) ListCampaigns() (map[int]*domain.Campaign, error) {

	campaigns, err := c.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return campaigns, nil

}

func (c *CampaignService) DonateToCampaign(id int, amount float64) (*domain.Campaign, error) {

	if amount <= 200 {
		return nil, errors.New("invalid donation amount")
	}

	domainCampaign, err := c.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = domainCampaign.AddDonation(amount)
	if err != nil {
		return nil, err
	}

	campaign, err := c.repo.Update(domainCampaign)
	if err != nil {
		return nil, err
	}
	return campaign, err

}

func (c *CampaignService) UpdateCampaignDetails(id int, newDescription string, newStatus domain.CampaignStatus) (*domain.Campaign, error) {

	domainCampaign, err := c.repo.FindByID(id)
	if err != nil {
		return nil, err

	}

	if newDescription == "" && newStatus == "" {
		return nil, errors.New("description and status cannot both be empty")
	}

	if len(newDescription) > 200 {
		return nil, errors.New("description exceeds maximum length of 200 characters")
	}

	domainCampaign.Description = newDescription
	domainCampaign.Status = newStatus
	campaign, err := c.repo.Update(domainCampaign)
	if err != nil {
		return nil, err
	}

	return campaign, nil

}
