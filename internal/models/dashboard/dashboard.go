package dashboard

type DashboardStats struct {
	// Aircraft Stats
	TotalAircrafts       int64 `json:"total_aircrafts"`
	ActiveAircrafts      int64 `json:"active_aircrafts"`
	MaintenanceAircrafts int64 `json:"maintenance_aircrafts"`

	// Crew Stats
	TotalCrews  int64 `json:"total_crews"`
	ActiveCrews int64 `json:"active_crews"`

	// Route & Airport Stats
	TotalRoutes   int64 `json:"total_routes"`
	ActiveRoutes  int64 `json:"active_routes"`
	TotalAirports int64 `json:"total_airports"`
}
