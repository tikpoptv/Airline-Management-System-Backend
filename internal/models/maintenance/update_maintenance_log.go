package maintenance

type UpdateMaintenanceLogRequest struct {
	DateOfMaintenance   *string `json:"date_of_maintenance"`
	Details             *string `json:"details"`
	MaintenanceLocation *string `json:"maintenance_location"`
	Status              *string `json:"status"` // optional: "Pending", "In Progress", "Completed", "Cancelled"
	AssignedTo          *uint   `json:"assigned_to"`
}
	