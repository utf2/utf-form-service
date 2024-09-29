package transfer

import (
	"time"

	"github.com/google/uuid"
)

type FormSearchByTeacherIDRequest struct {
	TeacherID uuid.UUID
}

type FormSearchByTeacherIDResponse struct {
	FormsCreatedByTeacher []CreatedFormDTO
}

type CreatedFormDTO struct {
	ID          uuid.UUID
	Name        string
	Description string
	CreateDate  time.Time
	StartDate   time.Time
	EndDate     time.Time
}

type FormSearchByIDRequest struct {
	FormID uuid.UUID
}

type FormSearchByIDResponse struct {
	Form FormDTO
}

type FormDTO struct {
	ID          uuid.UUID
	Name        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	TeacherID   uuid.UUID
	Questions   []QuestionDTO
}

type QuestionDTO struct {
	ID             uuid.UUID
	Text           string
	OrderNumber    uint
	CompletionTime time.Duration
	Answers        []AnswerDTO
}

type AnswerDTO struct {
	ID      uuid.UUID
	Text    string
	Correct bool
}
