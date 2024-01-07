package app

import (
	blittypes "blit/x/blit/types"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func ensure() {

	out, runErr := runBlitVmScript("ensure_blitvm.py")

	if runErr != nil {
		blitvmPath := viper.GetString(blittypes.FlagBlitVMPath)

		fmt.Println("Error running ensure_blitvm:", runErr, out)
		// if blitvmPath is relative, it's relative to the current working directory
		fmt.Printf("%s: %s\n", blittypes.FlagBlitVMPath, blitvmPath)
		if !filepath.IsAbs(blitvmPath) {
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Println("Error getting current working directory.", err)
			} else {

				fmt.Printf("current working directory: %s\nnormalized: %s\n", cwd, filepath.Join(cwd, blitvmPath))
				fmt.Printf("If you want to use blitd globally, you MUST to set %s to an absolute path. Currently it is a relative path.\n", blittypes.FlagBlitVMPath)
			}
		}
		// print multi line help
		fmt.Println(`
Assuming you checked out and built it in your home directory.
Set as a runtime flag:
----
$ blitd start --blit.blitvm_path /home/[username]/blitchain/blitvm
----

or set in ~/.blit/config/app.toml:
...
[blit]
  blitvm_path = "/home/[username]/blitchain/blitvm"
...

Or use an environment variable:
...
BLITD_BLIT_BLITVM_PATH=/home/[username]/blitchain/blitvm
...`)

		os.Exit(42)
	}

}
