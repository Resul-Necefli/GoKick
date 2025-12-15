package domain

import "errors"

var (
	ErrNotFound  = errors.New("campaign not found")
	ErrDuplicate = errors.New("campaign with given ID already exists")
)

// var (
// 	ErrInvalidAmount = errors.New("invalid donation amount")
// )

// var (
// 	ErrDescriptionTooLong = errors.New("description exceeds maximum length of 200 characters")
// )
// var (
// 	ErrNegativeID = errors.New("ID value cannot be negative")
// )
// var (
// 	ErrEmptyUpdateFields = errors.New("description and status cannot both be empty")
// )

// var (
// 	ErrCampaignNotActive  = errors.New("campaign is not active")
// 	ErrCampaignFinished   = errors.New("campaign already finished")
// )
