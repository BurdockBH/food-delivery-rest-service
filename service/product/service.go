package product

import (
	"encoding/json"
	product "github.com/BurdockBH/food-delivery-rest-service/db/product"
	"github.com/BurdockBH/food-delivery-rest-service/router/helper"
	"github.com/BurdockBH/food-delivery-rest-service/statusCodes"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"log"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p viewmodels.Product
	err := json.NewDecoder(r.Body).Decode(&p)
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

	err = p.ValidateProduct()
	if err != nil {
		log.Println("Failed to validate request body: ", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToValidateProduct,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToValidateProduct],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = product.CreateProduct(&p, email)
	if err != nil {
		log.Println("Failed to create product: ", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToCreateProduct,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToCreateProduct] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(viewmodels.BaseResponse{
		StatusCode: statusCodes.SuccesfullyCreatedProduct,
		Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyCreatedProduct] + ":" + p.Name,
	})
	log.Printf("Product %s created successfully", p.Name)
	helper.BaseResponse(w, response, http.StatusOK)
}
