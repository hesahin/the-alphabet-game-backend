package schema

import "the-alphabet-game-backend/api/datamodels"

type QuestionsResponse struct {
	Success   bool                   `json:"success"`
	Questions datamodels.QuestionSet `json:"questions"`
	Message   string                 `json:"message"`
}
