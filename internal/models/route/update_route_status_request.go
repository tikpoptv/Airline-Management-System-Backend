package route

type UpdateRouteStatusRequest struct {
	Status string `json:"status" binding:"required"`
}
