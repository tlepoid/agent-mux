package common

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

// AgentStatus is the aggregated conversation status for a workspace,
// derived from its tabs and used for dashboard display.
type AgentStatus int

const (
	AgentStatusIdle     AgentStatus = iota // ○ gray  - no agent or disconnected
	AgentStatusRunning                     // ● green  - actively working
	AgentStatusWaiting                     // ◐ yellow - needs user input
	AgentStatusError                       // ✕ red    - something went wrong
	AgentStatusComplete                    // ✓ info   - marked complete, awaiting review
)

// AgentStatusIcon returns the display icon for the given status.
func AgentStatusIcon(s AgentStatus) string {
	switch s {
	case AgentStatusRunning:
		return Icons.Running
	case AgentStatusWaiting:
		return Icons.Waiting
	case AgentStatusError:
		return Icons.Error
	case AgentStatusComplete:
		return Icons.Complete
	default:
		return Icons.Idle
	}
}

// Agent status colors remain constant across themes for consistent recognition.
var (
	colorStatusRunning  = lipgloss.Color("#98C379") // green
	colorStatusWaiting  = lipgloss.Color("#E5C07B") // yellow
	colorStatusError    = lipgloss.Color("#E06C75") // red
	colorStatusComplete = lipgloss.Color("#61AFEF") // blue
	colorStatusIdle     = lipgloss.Color("#5C6370") // gray
)

// AgentStatusColor returns the display color for the given status.
// Colors are fixed across themes so status indicators are always recognizable.
func AgentStatusColor(s AgentStatus) color.Color {
	switch s {
	case AgentStatusRunning:
		return colorStatusRunning
	case AgentStatusWaiting:
		return colorStatusWaiting
	case AgentStatusError:
		return colorStatusError
	case AgentStatusComplete:
		return colorStatusComplete
	default:
		return colorStatusIdle
	}
}
