package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) foo() {
	// Determine the command to launch the default browser on macOS
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", "https://www.google.com")
	default:
		fmt.Println("Unsupported platform")
		return
	}

	// Launch the default browser
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error launching browser:", err)
	}
}

func (a *App) LaunchChrome(url string) string {
	var cmd *exec.Cmd

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Windows")
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://www.google.com")
	case "linux":
		fmt.Println("Linux")
		cmd = exec.Command("xdg-open", url)
	case "darwin":
		fmt.Println("macOS")
		cmd = exec.Command("open", "https://www.google.com")
	default:
		fmt.Printf("Unknown operating system: %v\n", os)
	}

	err := cmd.Start()
	if err != nil {
		return fmt.Sprintf("failed to start Google Chrome: %v", err)
	}
	return ""
}

func fetchFileFromGitlab(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch file from gitlab: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
