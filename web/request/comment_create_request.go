package request

type CommentCreateRequest struct {
	Description string `json:"description" validate:"required"`
}
