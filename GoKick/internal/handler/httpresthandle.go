package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Resul-Necefli/GoKick/internal/domain"
	"github.com/Resul-Necefli/GoKick/internal/service"
)

type CampaignHandle struct {
	service *service.CampaignService
}

func NewCampaignHandler(s *service.CampaignService) *CampaignHandle {
	return &CampaignHandle{service: s}
}

func (c *CampaignHandle) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign domain.Campaign

	err := json.NewDecoder(r.Body).Decode(&campaign)
	if err != nil {
		log.Println("[Decode] error", err)
		http.Error(w, "json decode error", http.StatusBadRequest)
		return
	}

	err = c.service.CreateCampaign(&campaign)
	if err != nil {

		log.Println("[CreateCampaign] error", err)
		http.Error(w, "create campaign error", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(campaign)

}

// GetCampaignByID handles HTTP GET requests to retrieve a campaign by its numeric ID.
// The handler expects the request URL path to be in the form "/listID/{id}" and parses
func (c *CampaignHandle) GetCampaignByID(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	if id == "" {
		log.Println("[GetCampaignByID] error : id is empty")
		http.Error(w, "id is empty", http.StatusBadRequest)
		return
	}

	convertID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("[strcnov] error : not type convert")
		http.Error(w, "not type convert", http.StatusBadRequest)
		return
	}

	campaign, err := c.service.GetCampaignByID(convertID)
	if err != nil {
		log.Println("[service GetCampaignByID] error :", err)
		http.Error(w, "service error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(campaign)

}

func (c *CampaignHandle) ListCampaigns(w http.ResponseWriter, r *http.Request) {
	// yalniz ozune aid olan lar list edilmelidir heleki o mentiqde islemir amma bunu hisseler duzeldilmelidir
	campaigns, err := c.service.ListCampaigns()
	if err != nil {
		log.Println("[service ListCampaigns] error :", err)
		http.Error(w, "service error", http.StatusInternalServerError)
		return
	}

	sliceDomain := make([]*domain.Campaign, 0, len(campaigns))

	for _, v := range campaigns {

		sliceDomain = append(sliceDomain, v)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sliceDomain)

}

func (c *CampaignHandle) UpdateDetails(w http.ResponseWriter, r *http.Request) {

	type DescripStatusReq struct {
		Description string                `json:"description"`
		Status      domain.CampaignStatus `json:"status"`
	}

	var data DescripStatusReq
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		log.Println("[UpdateDetails DescripStatusReq ] not is decode ", err)
		http.Error(w, "internal server error", http.StatusBadRequest)
		return
	}

	id := r.PathValue("id")

	itID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("[strconv] error : not is  string convert")
		http.Error(w, "not is convert string", http.StatusBadRequest)
	}

	domainObject, err := c.service.UpdateCampaignDetails(itID, data.Description, data.Status)

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			log.Println("[service UpdateCampaignDetails] error : campaign not found")
			http.Error(w, "campaign not found", http.StatusNotFound)
			return
		}
		log.Println("[service UpdateCampaignDetails] error :", err)
		http.Error(w, "service error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(domainObject)

}

func (h *CampaignHandle) DonateHandler(w http.ResponseWriter, r *http.Request) {

	type PaymentRequest struct {
		Amount float64 `json:"amount"`
	}

	var payment PaymentRequest

	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		log.Println("[DonateHandler Decoder] error :", err)
		http.Error(w, "not is decode", http.StatusBadRequest)
		return
	}

	id := r.PathValue("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("[strconv] error : not is  string convert")
		http.Error(w, "not is convert string", http.StatusBadRequest)
	}

	_, err = h.service.DonateToCampaign(intID, payment.Amount)

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			log.Println("[service DonateToCampaign] error : campaign not found")
			http.Error(w, "campaign not found", http.StatusNotFound)
			return
		}

		log.Println("[service DonateToCampaign] error :", err)
		http.Error(w, "service error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
