// Package vterm implements a virtual terminal emulator (VTerm) that parses
// ANSI/VT100 escape sequences and maintains an in-memory cell grid. It
// supports the alternate screen buffer, scrollback history, cursor tracking,
// wide characters, and selection. The rendered output is consumed by the
// compositor layer for efficient TUI rendering without full-screen redraws.
package vterm
