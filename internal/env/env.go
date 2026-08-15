package env

import (
	"os"
	"os/exec"
)

type ClipboardCapabilities struct {
    Wayland bool
    X11     bool
}

type Environment struct {
    Clipboard ClipboardCapabilities
}

func Collect() Environment {
	env := Environment{}
	_, waylandClipboardExists := os.LookupEnv("WAYLAND_DISPLAY")
	_, x11ClipboardExists := os.LookupEnv("DISPLAY")

	if waylandClipboardExists {
		if _, err := exec.LookPath("wl-copy"); err == nil {
        env.Clipboard.Wayland = true
    	}
	}
	
	if x11ClipboardExists {
		if _, err := exec.LookPath("xclip"); err == nil {
        env.Clipboard.X11 = true
    	}
	}
	
	return env
}