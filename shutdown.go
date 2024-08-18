package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func shutdown() {
	var cmd *exec.Cmd

	switch os := runtime.GOOS; os {
	case "windows":
		cmd = exec.Command("shutdown", "/s", "/t", "0")
	case "linux":
		cmd = exec.Command("shutdown", "now")
	case "darwin":
		cmd = exec.Command("sudo", "shutdown", "-h", "now")
	default:
		fmt.Printf("Unsupported OS: %s\n", os)
		return
	}

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Shutdown command failed: %v\n", err)
	} else {
		fmt.Println("Shutdown command executed successfully")
	}
}

func main() {
	shutdown()
}
