package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func ensure() {

	out, runErr := runBlitVmScript("ensure_blitvm.py")

	if runErr != nil {
		blitvmPath := viper.GetString("blit.experimental_blitvm_path")

		fmt.Println("Error running ensure_blitvm:", runErr, out)
		fmt.Printf("blit.experimental_blitvm_path: %s\n", blitvmPath)
		// if blitvmPath is relative, it's relative to the current working directory
		if !filepath.IsAbs(blitvmPath) {
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Println("Error getting current working directory.", err)
			} else {
				fmt.Printf("current working directory: %s\nnormalized: %s\n", cwd, filepath.Join(cwd, blitvmPath))
			}
		}

		os.Exit(42)
	}

}
