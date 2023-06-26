package product

import (
	"encoding/json"
	product "github.com/BurdockBH/food-delivery-rest-service/db/product"
	"github.com/BurdockBH/food-delivery-rest-service/router/helper"
	"github.com/BurdockBH/food-delivery-rest-service/statusCodes"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"log"
	"net/http"
	"strings"
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

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		log.Println("Token not found")
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.TokenNotFound,
			Message:    statusCodes.StatusCodes[statusCodes.TokenNotFound],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	claims, err := helper.ValidateToken(tokenString)
	if err != nil {
		log.Println("Token validation failed:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.TokenValidationFailed,
			Message:    statusCodes.StatusCodes[statusCodes.TokenValidationFailed],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

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

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
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

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		log.Println("Token not found")
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.TokenNotFound,
			Message:    statusCodes.StatusCodes[statusCodes.TokenNotFound],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	_, err = helper.ValidateToken(tokenString)
	if err != nil {
		log.Println("Token validation failed:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.TokenValidationFailed,
			Message:    statusCodes.StatusCodes[statusCodes.TokenValidationFailed],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = product.DeleteProduct(&p)
	if err != nil {
		log.Println("Failed to create product: ", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToDeleteProduct,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToDeleteProduct] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(viewmodels.BaseResponse{
		StatusCode: statusCodes.SuccesfullyDeletedProduct,
		Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyDeletedProduct] + ":" + p.Name,
	})
	log.Printf("Product %s deleted successfully", p.Name)
	helper.BaseResponse(w, response, http.StatusOK)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
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

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		log.Println("Token not found")
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.TokenNotFound,
			Message:    statusCodes.StatusCodes[statusCodes.TokenNotFound],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	_, err = helper.ValidateToken(tokenString)
	if err != nil {
		log.Println("Token validation failed:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.TokenValidationFailed,
			Message:    statusCodes.StatusCodes[statusCodes.TokenValidationFailed],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = product.EditProduct(&p)
	if err != nil {
		log.Println("Failed to edit product: ", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToUpdateProduct,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToUpdateProduct] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(viewmodels.BaseResponse{
		StatusCode: statusCodes.SuccesfullyUpdatedProduct,
		Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyUpdatedProduct] + ":" + p.Name,
	})
	log.Printf("Product %s deleted successfully", p.Name)
	helper.BaseResponse(w, response, http.StatusOK)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
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

	products, err := product.GetProducts(&p)
	if err != nil {
		log.Println("Failed to get products: ", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToGetProducts,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToGetProducts] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(viewmodels.ProductList{
		Products: products,
		BaseResponse: viewmodels.BaseResponse{
			StatusCode: statusCodes.SuccesfullyFetchedProducts,
			Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyFetchedProducts] + ":" + p.Name,
		},
	})
	log.Printf("Products fetched successfully")
	helper.BaseResponse(w, response, http.StatusOK)
}
