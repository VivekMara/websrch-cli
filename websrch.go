package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	// "net"
	// "os"
	// "os/exec"
	"runtime"
	// "strings"
)

func main() {
	fmt.Println(runtime.GOOS)
	url := "https://www.google.com/search?q="
	query := strings.Join(os.Args[1:], " ")
	srchURL := url + query

	err := openBrowser(srchURL)
    if err != nil {
        fmt.Println("Error opening browser:", err)
    }
	
}

func openBrowser(url string) error {
    var err error

    switch runtime.GOOS {
    case "linux":
        err = exec.Command("xdg-open", url).Start()
    case "windows":
        err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
    case "darwin":
        err = exec.Command("open", url).Start()
    default:
        err = fmt.Errorf("unsupported platform")
    }
    return err
}