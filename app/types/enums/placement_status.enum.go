package enums

type PlacementStatus string

var (
	PlacementOngoingStatus PlacementStatus = "ongoing"
	PlacementSuspendStatus PlacementStatus = "suspend"
	PlacementCancelStatus  PlacementStatus = "cancel"
	PlacementEndStatus     PlacementStatus = "end"
)
