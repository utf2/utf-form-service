package transfer

import (
	"time"

	"github.com/google/uuid"
)

type FormCreateRequest struct {
	Name        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	TeacherID   uuid.UUID
	Questions   []QuestionCreateRequest
}

type QuestionCreateRequest struct {
	Text           string
	OrderNumber    uint
	CompletionTime time.Duration
	Answers        []AnswerCreateRequest
}

type AnswerCreateRequest struct {
	Text    string
	Correct bool
}

type FormCreateResponse struct {
	Success bool
	FormID  uuid.UUID
}
