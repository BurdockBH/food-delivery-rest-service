package viewmodels

import "errors"

type ItemIdRequest struct {
	Id int64 `json:"id"`
}

func (i *ItemIdRequest) ValidateItemIdRequest() error {
	if i.Id < 0 {
		return errors.New("Id must be positive")
	}
	return nil
}
