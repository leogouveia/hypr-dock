package ipc

import (
	"os"
	"strconv"
	"strings"
)

// luaCommand is opt-in so older Hyprland sessions retain their original IPC.
// Set HYPR_DOCK_LUA=1 in the launcher for a Hyprland Lua session.
func luaCommand(cmd string) string {
	if os.Getenv("HYPR_DOCK_LUA") != "1" {
		return cmd
	}
	if address, ok := strings.CutPrefix(cmd, "dispatch focuswindow address:"); ok {
		selector := strconv.Quote("address:" + address)
		// Focusing and raising are separate operations for floating windows.
		// Suppress pointer warps only inside this synchronous action; always
		// restore the user's preference, including when either action errors.
		return "dispatch (function() local old=hl.get_config(\"cursor.no_warps\"); " +
			"hl.config({cursor={no_warps=true}}); " +
			"local ok,result=pcall(function() " +
			"hl.dispatch(hl.dsp.focus({window=" + selector + "})); " +
			"hl.dispatch(hl.dsp.window.alter_zorder({window=" + selector + ",mode=\"top\"})); " +
			"end); hl.config({cursor={no_warps=old}}); " +
			"if not ok then error(result) end; return result end)"
	}
	if address, ok := strings.CutPrefix(cmd, "dispatch closewindow address:"); ok {
		return "dispatch hl.dsp.window.close({window=" + strconv.Quote("address:"+address) + "})"
	}
	return cmd
}
