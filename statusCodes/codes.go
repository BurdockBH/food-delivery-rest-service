package statusCodes

const (
	UserDoesNotExist              = 1
	FailedToDecodeRequestBody     = 2
	FailedToValidateUser          = 3
	FailedToCreateUser            = 4
	FailedToUpdateUser            = 5
	FailedToValidateLogin         = 6
	FailedToLoginUser             = 7
	FailedToGenerateToken         = 8
	FailedToMarshalJSON           = 9
	FailedToWriteResponse         = 10
	TokenNotFound                 = 11
	TokenValidationFailed         = 12
	InvalidClaims                 = 13
	FailedToDeleteUser            = 14
	FailedToFetchUsers            = 15
	SuccesfullyCreatedUser        = 16
	SuccesfullyDeletedUser        = 17
	SuccesfullyFetchedUsers       = 18
	SuccesfullyUpdatedUser        = 19
	SuccesfullyLoggedInUser       = 20
	FailedToValidateFoodVenue     = 21
	SuccesfullyCreatedFoodVenue   = 22
	FailedToCreateFoodVenue       = 23
	FailedToDeleteFoodVenue       = 24
	SuccesfullyDeletedFoodVenue   = 25
	FailedToFetchFoodVenues       = 26
	SuccesfullyFetchedFoodVenues  = 27
	FailedToValidateProduct       = 28
	FailedToCreateProduct         = 29
	SuccesfullyCreatedProduct     = 30
	FailedToDeleteProduct         = 31
	SuccesfullyDeletedProduct     = 32
	FailedToUpdateProduct         = 33
	SuccesfullyUpdatedProduct     = 34
	FailedToGetProducts           = 35
	SuccesfullyFetchedProducts    = 36
	FailedToValidateItemIdRequest = 37
)

var StatusCodes = map[int64]string{
	UserDoesNotExist:              "User does not exist",
	FailedToDecodeRequestBody:     "Failed to decode request body",
	FailedToValidateUser:          "Failed to validate user",
	FailedToCreateUser:            "Failed to create user",
	FailedToUpdateUser:            "Failed to update user",
	FailedToValidateLogin:         "Failed to validate login",
	FailedToLoginUser:             "Failed to login user",
	FailedToGenerateToken:         "Failed to generate token",
	FailedToMarshalJSON:           "Failed to marshal json",
	FailedToWriteResponse:         "Failed to write response",
	TokenNotFound:                 "Token not found",
	TokenValidationFailed:         "Token validation failed",
	InvalidClaims:                 "Invalid claims",
	FailedToDeleteUser:            "Failed to delete user",
	FailedToFetchUsers:            "Failed to fetch users",
	SuccesfullyCreatedUser:        "User created successfully!",
	SuccesfullyDeletedUser:        "User deleted successfully!",
	SuccesfullyFetchedUsers:       "Users fetched successfully!",
	SuccesfullyUpdatedUser:        "User updated successfully!",
	SuccesfullyLoggedInUser:       "User logged in successfully!",
	FailedToValidateFoodVenue:     "Failed to validate food venue",
	SuccesfullyCreatedFoodVenue:   "Food venue created successfully!",
	FailedToCreateFoodVenue:       "Failed to create food venue",
	FailedToDeleteFoodVenue:       "Failed to delete food venue",
	SuccesfullyDeletedFoodVenue:   "Food venue deleted successfully!",
	FailedToFetchFoodVenues:       "Failed to fetch food venues",
	SuccesfullyFetchedFoodVenues:  "Food venues fetched successfully!",
	FailedToValidateProduct:       "Failed to validate product",
	FailedToCreateProduct:         "Failed to create product",
	SuccesfullyCreatedProduct:     "Product created successfully!",
	FailedToDeleteProduct:         "Failed to delete product",
	SuccesfullyDeletedProduct:     "Product deleted successfully!",
	FailedToUpdateProduct:         "Failed to update product",
	SuccesfullyUpdatedProduct:     "Product updated successfully!",
	FailedToGetProducts:           "Failed to get products",
	SuccesfullyFetchedProducts:    "Products fetched successfully!",
	FailedToValidateItemIdRequest: "Failed to validate item id request",
}
