package models

const (
	StatusInactive  = 0
	StatusActive    = 1
	StatusSuspended = 8
)

var StatusLabels = map[int]string{
	StatusInactive:  "Inactive",
	StatusActive:    "Active",
	StatusSuspended: "Suspended",
}

var StatusColors = map[int]string{
	StatusInactive:  "red",
	StatusActive:    "green",
	StatusSuspended: "yellow",
}

var StatusBadges = map[int]string{
	StatusInactive:  "Deactivated",
	StatusActive:    "Active",
	StatusSuspended: "Suspended",
}


// MapStatus returns the label and color for a given status.
// If the status is nil or unknown, it returns "Unknown" and "gray".
func MapStatus(status *int) (label string, color string, badge string) {
	if status == nil {
		return "Unknown", "gray", "Unknown"
	}
	label, ok1 := StatusLabels[*status]
	color, ok2 := StatusColors[*status]
	badge, ok3 := StatusBadges[*status]
	if !ok1 {
		label = "Unknown"
	}
	if !ok2 {
		color = "gray"
	}
	if !ok3 {
		badge = "Unknown"
	}
	return label, color, badge
}