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
			return errors.New("could not copy to Wayland clipboard: backend failed")
		}
		return nil
	}
	if clipboard.X11 {
		cmd := exec.Command("xclip", "-selection", "clipboard")
		cmd.Stdin = strings.NewReader(data)
		err := cmd.Run()
		if err != nil {
			return errors.New("could not copy to X11 clipboard: backend failed")
		}
		return nil
	}

	return errors.New("Cannot find any clipboard backend")
}