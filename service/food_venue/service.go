package food_venue

import (
	"encoding/json"
	"github.com/BurdockBH/food-delivery-rest-service/db/food_venue"
	"github.com/BurdockBH/food-delivery-rest-service/router/helper"
	"github.com/BurdockBH/food-delivery-rest-service/statusCodes"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"log"
	"net/http"
)

func CreateFoodVenue(w http.ResponseWriter, r *http.Request) {
	var foodVenue viewmodels.FoodVenue
	err := json.NewDecoder(r.Body).Decode(&foodVenue)
	if err != nil {
		log.Println("Failed to decode request body: ", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToDecodeRequestBody,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToDecodeRequestBody],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	claims := *helper.CheckToken(&w, r)

	email := claims["email"].(string)

	err = foodVenue.ValidateFoodVenue()
	if err != nil {
		log.Println("Failed to validate request body: ", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToValidateFoodVenue,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToValidateFoodVenue],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = food_venue.CreateFoodVenue(&foodVenue, email)
	if err != nil {
		log.Println("Failed to create food venue: ", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToCreateFoodVenue,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToCreateFoodVenue] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(viewmodels.BaseResponse{
		StatusCode: statusCodes.SuccesfullyCreatedFoodVenue,
		Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyCreatedFoodVenue] + ":" + foodVenue.Name,
	})
	log.Printf("Food venue created: %s", foodVenue.Name)
	helper.BaseResponse(w, response, http.StatusOK)
}

func DeleteFoodVenue(w http.ResponseWriter, r *http.Request) {
	var foodVenue viewmodels.FoodVenue
	err := json.NewDecoder(r.Body).Decode(&foodVenue)
	if err != nil {
		log.Println("Failed to decode request body: ", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToDecodeRequestBody,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToDecodeRequestBody],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	_ = helper.CheckToken(&w, r)

	err = food_venue.DeleteFoodVenue(&foodVenue)
	if err != nil {
		log.Println("Failed to delete food venue: ", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToDeleteFoodVenue,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToDeleteFoodVenue] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(viewmodels.BaseResponse{
		StatusCode: statusCodes.SuccesfullyDeletedFoodVenue,
		Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyDeletedFoodVenue] + ":" + foodVenue.Name,
	})

	log.Printf("Food venue deleted: %s", foodVenue.Name)
	helper.BaseResponse(w, response, http.StatusOK)
}

func GetFoodVenues(w http.ResponseWriter, r *http.Request) {
	var foodVenue viewmodels.FoodVenue
	err := json.NewDecoder(r.Body).Decode(&foodVenue)
	if err != nil {
		log.Println("Failed to decode request body: ", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToDecodeRequestBody,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToDecodeRequestBody],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	foodVenues, err := food_venue.GetVenues(&foodVenue)
	if err != nil {
		log.Println("Failed to get food venues: ", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToFetchFoodVenues,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToFetchFoodVenues] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(viewmodels.FoodVenueList{
		BaseResponse: viewmodels.BaseResponse{
			StatusCode: statusCodes.SuccesfullyFetchedFoodVenues,
			Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyFetchedFoodVenues],
		},
		FoodVenues: foodVenues,
	})

	log.Printf("Food venues fetched: %d", len(foodVenues))
	helper.BaseResponse(w, response, http.StatusOK)
}
