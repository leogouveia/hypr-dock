package ipc

import (
	"encoding/json"
	"fmt"
	"image/png"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// This opt-in test operates only on a disposable fixture window. It exercises
// the actual IPC entry point used by icon and thumbnail clicks, not a copied
// version of the Lua expression. Run with HYPR_DOCK_TEST_TARGET=0x... .
func TestLuaFocusIntegration(t *testing.T) {
	address := os.Getenv("HYPR_DOCK_TEST_TARGET")
	if address == "" {
		t.Skip("requires a disposable Hyprland fixture window")
	}
	window, err := SearchClientByAddress(address)
	if err != nil || window.Class != "hypr-dock-focus-test.py" || window.Title != "Teste temporário hypr-dock" {
		t.Fatal("target must be a disposable hypr-dock-focus-test.py fixture", err)
	}
	t.Setenv("HYPR_DOCK_LUA", "1")
	read := func(cmd string) map[string]interface{} {
		t.Helper()
		data, err := Hyprctl(cmd)
		if err != nil {
			t.Fatal(err)
		}
		var value map[string]interface{}
		if err := json.Unmarshal(data, &value); err != nil {
			t.Fatal(err)
		}
		return value
	}
	before := read("j/cursorpos")
	option := read("j/getoption cursor:no_warps")
	response, err := Hyprctl("dispatch focuswindow address:" + address)
	if err != nil || strings.TrimSpace(string(response)) != "ok" {
		t.Fatalf("focus dispatch failed: %s %v", response, err)
	}
	active, err := GetActiveWindow()
	if err != nil || active.Address != address {
		t.Fatal("wrong active window", active, err)
	}
	after := read("j/cursorpos")
	if before["x"] != after["x"] || before["y"] != after["y"] {
		t.Fatalf("pointer moved: %v -> %v", before, after)
	}
	if option["bool"] != read("j/getoption cursor:no_warps")["bool"] {
		t.Fatal("global cursor preference was not restored")
	}
	// The fixture target is red and is initially covered by a blue window.
	// Pixel verification catches the "focused, but still behind" regression.
	time.Sleep(200 * time.Millisecond)
	path := filepath.Join(t.TempDir(), "focused.png")
	geometry := fmt.Sprintf("%d,%d 40x40", window.At[0]+70, window.At[1]+70)
	if data, err := exec.Command("grim", "-g", geometry, path).CombinedOutput(); err != nil {
		t.Fatalf("cannot capture fixture: %s %v", data, err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		t.Fatal(err)
	}
	r, _, b, _ := img.At(20, 20).RGBA()
	if r < 60000 || b > 5000 {
		t.Fatalf("target not raised: r=%d b=%d", r, b)
	}
}
