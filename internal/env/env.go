package env

import (
	"os"
	"os/exec"
)

type Environment struct {
    WaylandClipboard bool
    X11Clipboard    bool
}

func Collect() Environment {
	env := Environment{}
	_, waylandClipboardExists := os.LookupEnv("WAYLAND_DISPLAY")
	_, x11ClipboardExists := os.LookupEnv("DISPLAY")

	if waylandClipboardExists {
		if _, err := exec.LookPath("wl-copy"); err == nil {
        env.WaylandClipboard = true
    	}
	}
	
	if x11ClipboardExists {
		if _, err := exec.LookPath("xclip"); err == nil {
        env.X11Clipboard = true
    	}
	}
	
	return env
}