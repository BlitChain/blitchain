package app

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func ensure() {
	// Verify blit-python is installed

	blitvmPath := os.Getenv("BLITVM_PATH")
	if blitvmPath == "" {
		blitvmPath = "./blitvm/"
	}

	out, runErr := exec.Command(
		"python3", filepath.Join(blitvmPath, "ensure_blitvm.py"),
	).CombinedOutput()

	if runErr != nil {
		fmt.Println("Error running ensure_blitvm.", runErr, string(out))
		fmt.Printf("BLITVM_PATH: %s\n", blitvmPath)
		os.Exit(42)
	}

}
