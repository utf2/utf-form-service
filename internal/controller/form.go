package controller

import transfer "github.com/utf2/utf-form-service/internal/controller/model"

type FormControllerAPI interface {
	Create(transfer.FormCreateRequest) transfer.FormCreateResponse
	SearchByTeacherID(transfer.FormSearchByTeacherIDRequest) transfer.FormSearchByTeacherIDResponse
	SearchByIDs(transfer.FormSearchByIDsRequest) transfer.FormSearchByIDsResponse
}
