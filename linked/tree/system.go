package tree

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func isDotCommandInstalled() bool {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux", "darwin": // Linux or macOS
		cmd = exec.Command("dot", "-V")
	case "windows": // Windows
		cmd = exec.Command("where", "dot")
	default:
		fmt.Println("Unsupported OS")
		return false
	}

	// Ensure the PATH is set correctly (if needed)
	cmd.Env = append(os.Environ(), "PATH=/usr/bin:"+os.Getenv("PATH"))

	err := cmd.Run()
	return err == nil
}
