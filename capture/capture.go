/*Package capture is an interface to the Raspberry Pi camera.

Currently, it uses the raspistill binary to pull an image with a timeout of 1
second.  This works okay if you don't want images more often than about every 2
seconds.  Future improvemets might use the MMAL library (via
https://github.com/djthorpe/gopi) to get more stable images, at a higher
framerate.
*/
package capture

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"syscall"
)

var raspistillPath = "/usr/bin/raspistill"

// CameraOptions provides a way to set options to the camera.
type CameraOptions struct {
	hf bool // flip the image horizontally
	vf bool // flip the image vertically
}

// GetImage takes options as a struct, calls raspistill, and returns an
// io.Writer with the image data from the camera.
func GetImage(opts *CameraOptions) ([]byte, error) {
	args := []string{"-o", "-", "--nopreview", "--timeout", "1"}

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
					return []byte{}, fmt.Errorf("please run raspi-config and enable the camera")
				}
			}
		}
		return []byte{}, err
	}

	return imgBytes, nil
}
