package maintenance

type UpdateTaskStatusRequest struct {
	Status  string  `json:"status" validate:"required,oneof=Pending In Progress Completed Cancelled"`
	Details *string `json:"details"`
}
