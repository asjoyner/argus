package capture

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"syscall"
)

var raspistillPath = "/usr/bin/raspistill"

type CameraOptions struct {
	hf bool // flip the image horizontally
	vf bool // flip the image vertically
}

// GetImage takes options as a struct, calls raspistill, and returns an
// io.Writer with the image data from the camera.
func GetImage(opts *CameraOptions) ([]byte, error) {
	args := []string{"-o", "-"}

	if opts != nil {
		if opts.hf {
			args = append(args, "-hf")
		}
		if opts.vf {
			args = append(args, "-vf")
		}
	}

	cmd := exec.Command(raspistillPath, args...)
	out, err := cmd.StdoutPipe()
	if err != nil {
		return []byte{}, err
	}

	err = cmd.Start()
	if err != nil {
		return []byte{}, err
	}

	imgBytes, err := ioutil.ReadAll(out)
	if err != nil {
		return []byte{}, err
	}

	if err := cmd.Wait(); err != nil {
		// The program has exited with an exit code != 0
		if exiterr, ok := err.(*exec.ExitError); ok {
			// Extracting an error code from cmd.Wait() is platform dependent, but
			// since this is raspistill-specific, that should be fine.  :)
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				if status.ExitStatus() == 70 {
					return []byte{}, fmt.Errorf("Please run raspi-config and enable the camera.")
				}
			}
		}
		return []byte{}, err
	}

	return imgBytes, nil
}
