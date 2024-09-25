package transfer

import (
	"time"

	"github.com/google/uuid"
)

type FormSearchByTeacherIDRequest struct {
	TeacherID uuid.UUID
}

type FormSearchByTeacherIDResponse struct {
	FormsCreatedByTeacher []FormDTO
}

// Show form for student when answering
// Show all forms assigned to student
type FormSearchByIDsRequest struct {
	FormIDs uuid.UUIDs
}

type FormSearchByIDsResponse struct {
	Forms []FormDTO
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
