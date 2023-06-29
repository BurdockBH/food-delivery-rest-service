package product

import (
	"encoding/json"
	product "github.com/BurdockBH/food-delivery-rest-service/db/product"
	"github.com/BurdockBH/food-delivery-rest-service/router/helper"
	"github.com/BurdockBH/food-delivery-rest-service/statusCodes"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"log"
	"net/http"
	"strconv"
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

	claims := helper.CheckToken(&w, r)
	if claims == nil {
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
	var id viewmodels.ItemIdRequest
	err := json.NewDecoder(r.Body).Decode(&id)
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

	err = id.ValidateItemIdRequest()

	if err != nil {
		log.Println("Failed to validate request body: ", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToValidateItemIdRequest,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToValidateItemIdRequest],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	claims := helper.CheckToken(&w, r)
	if claims == nil {
		return
	}

	err = product.DeleteProduct(id.Id)
	if err != nil {
		log.Println("Failed to delete product: ", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToDeleteProduct,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToDeleteProduct] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(viewmodels.BaseResponse{
		StatusCode: statusCodes.SuccesfullyDeletedProduct,
		Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyDeletedProduct] + ":" + string(id.Id),
	})
	log.Printf("Product with id %v deleted successfully", id)
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

	if c := helper.CheckToken(&w, r); c == nil {
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
	log.Printf("Product %s updated successfully", p.Name)
	helper.BaseResponse(w, response, http.StatusOK)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var venueId viewmodels.ItemIdRequest
	err := json.NewDecoder(r.Body).Decode(&venueId)
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

	err = venueId.ValidateItemIdRequest()

	if err != nil {
		log.Println("Failed to validate id request body: ", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToValidateItemIdRequest,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToValidateItemIdRequest],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	products, err := product.GetProducts(venueId.Id)
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
			Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyFetchedProducts] + ":" + string(venueId.Id),
		},
	})
	log.Printf("Products fetched successfully")
	helper.BaseResponse(w, response, http.StatusOK)
}

func OrderProduct(w http.ResponseWriter, r *http.Request) {
	var order viewmodels.Order
	err := json.NewDecoder(r.Body).Decode(&order)
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

	claims := helper.CheckToken(&w, r)
	if claims == nil {
		return
	}

	email := claims["email"].(string)

	err = order.ValidateOrder()
	if err != nil {
		log.Println("Failed to validate request body: ", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToValidateOrder,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToValidateOrder],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = product.OrderProduct(&order, email)
	if err != nil {
		log.Println("Failed to order product: ", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToOrderProduct,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToOrderProduct] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(viewmodels.BaseResponse{
		StatusCode: statusCodes.SuccesfullyOrderedProduct,
		Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyOrderedProduct] + " id: " + strconv.FormatInt(order.ProductID, 10),
	})

	log.Printf("Product %s ordered successfully", string(order.ID))
	helper.BaseResponse(w, response, http.StatusOK)

}
