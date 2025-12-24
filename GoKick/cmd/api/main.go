package main

import (
	"log"
	"net/http"

	"github.com/Resul-Necefli/GoKick/internal/domain"
	"github.com/Resul-Necefli/GoKick/internal/handler"
	"github.com/Resul-Necefli/GoKick/internal/repository"
	"github.com/Resul-Necefli/GoKick/internal/service"
)

func main() {

	mux := http.NewServeMux()

	repo := repository.InMemoryCampaignRepository{
		Data: make(map[int]*domain.Campaign, 1000),
	}
	service := service.NewCampaignService(&repo)
	handler := handler.NewCampaignHandler(service)

	mux.HandleFunc("POST /campaigns", handler.CreateCampaign)
	mux.HandleFunc("GET /campaigns", handler.ListCampaigns)
	mux.HandleFunc("GET /campaigns/{id}", handler.GetCampaignByID)
	mux.HandleFunc("PUT /campaigns/{id}", handler.UpdateDetails)
	mux.HandleFunc("POST /campaigns/{id}/donate", handler.DonateHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))

}
