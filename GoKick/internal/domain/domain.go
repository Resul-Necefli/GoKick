package domain

import (
	"errors"
)

type CampaignStatus string

const (
	CampaignStatusActive  CampaignStatus = "active"
	CampaignStatusFinshed CampaignStatus = "Finshed"
)

type Campaign struct {
	ID            int            `json:"id"`
	Description   string         `json:"description"`
	TargetAmount  float64        `json:"targetAmount"`
	CurrentAmount float64        `json:"currentAmount"`
	Status        CampaignStatus `json:"status"`
}

func (c *Campaign) AddDonation(amount float64) error {

	if c.Status != CampaignStatusActive {
		return errors.New("campaign is not active")

	}
	if c.IsFinished() {
		c.Status = CampaignStatusFinshed
		return errors.New("campaign already finished")
	}

	c.CurrentAmount += amount
	return nil
}

func (c *Campaign) IsFinished() bool {

	return c.CurrentAmount >= c.TargetAmount

}
