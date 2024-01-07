package app

import (
	blittypes "blit/x/blit/types"

	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/viper"
)

func runBlitVmScript(script string) (string, error) {

	blitvmPath := viper.GetString(blittypes.FlagBlitVMPath)
	pyenv_root := os.Getenv("PYENV_ROOT")
	if pyenv_root == "" {
		fmt.Println("PYENV_ROOT is not set. Please set it to the root of your pyenv installation.")
		os.Exit(41)
	}

	pythonExe := filepath.Join(pyenv_root, "versions", "blit-python", "bin", "python")

	out, runErr := exec.Command(pythonExe, filepath.Join(blitvmPath, script)).CombinedOutput()
	return string(out), runErr
}
