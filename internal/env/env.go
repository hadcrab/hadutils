package env

import (
	"os"
	"os/exec"
)

type Environment struct {
    Wayland bool
    X11     bool
}

func Collect() Environment {
	env := Environment{
		Wayland: false,
		X11: false,
	}
	_, waylandExists := os.LookupEnv("WAYLAND_DISPLAY")
	_, x11Exists := os.LookupEnv("DISPLAY")

	if waylandExists {
		if _, err := exec.LookPath("wl-copy"); err == nil {
        env.Wayland = true
    	}
	}
	
	if x11Exists {
		if _, err := exec.LookPath("xclip"); err == nil {
        env.X11 = true
    	}
	}
	
	return env
}