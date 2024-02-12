package request

type CommentUpdateRequest struct {
	Id          int    `json:"id" validate:"required"`
	Description string `json:"description" validate:"required"`
}
