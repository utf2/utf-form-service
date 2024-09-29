package controller

import transfer "github.com/utf2/utf-form-service/internal/controller/model"

type FormControllerAPI interface {
	Create(transfer.FormCreateRequest) transfer.FormCreateResponse
	SearchByTeacherID(transfer.FormSearchByTeacherIDRequest) transfer.FormSearchByTeacherIDResponse
	SearchByID(transfer.FormSearchByIDRequest) transfer.FormSearchByIDResponse
}
