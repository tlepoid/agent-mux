package common

import "image/color"

// AgentStatus is the aggregated conversation status for a workspace,
// derived from its tabs and used for dashboard display.
type AgentStatus int

const (
	AgentStatusIdle    AgentStatus = iota // ○ gray  - no agent or disconnected
	AgentStatusRunning                    // ● green  - actively working
	AgentStatusWaiting                    // ◐ yellow - needs user input
	AgentStatusError                      // ✕ red    - something went wrong
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
	default:
		return Icons.Idle
	}
}

// AgentStatusColor returns the display color for the given status.
func AgentStatusColor(s AgentStatus) color.Color {
	switch s {
	case AgentStatusRunning:
		return ColorSuccess()
	case AgentStatusWaiting:
		return ColorWarning()
	case AgentStatusError:
		return ColorError()
	default:
		return ColorMuted()
	}
}
