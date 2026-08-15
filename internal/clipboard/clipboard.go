package clipboard

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/hadcrab/hadutils/internal/env"
)

func CopyIn(data string, clipboard env.ClipboardCapabilities) error {
	if clipboard.Wayland {
		cmd := exec.Command("wl-copy")
		cmd.Stdin = strings.NewReader(data)
		err := cmd.Run()
		if err != nil {
			return errors.New("Cannot call Wayland clipboard backend")
		}
		return nil
	}
	if clipboard.X11 {
		cmd := exec.Command("xclip", "-selection", "clipboard")
		cmd.Stdin = strings.NewReader(data)
		err := cmd.Run()
		if err != nil {
			return errors.New("Cannot call X11 clipboard backend")
		}
		return nil
	}

	return errors.New("Cannot find any clipboard backend")
}