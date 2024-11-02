package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World! This is a simple Go web server.")
    })

    // Start the web server in a separate goroutine
    go func() {
        fmt.Println("Starting server on :8080...")
        if err := http.ListenAndServe(":8080", nil); err != nil {
            fmt.Println("Error starting server:", err)
        }
    }()

    // Open the default web browser
    openBrowser("http://localhost:8080")

    // Prevent the main function from exiting
    select {}
}

func openBrowser(url string) {
    var err error
    switch runtime.GOOS {
    case "darwin": // MacOS
        err = exec.Command("open", url).Start()
    case "linux":
        err = exec.Command("xdg-open", url).Start()
    case "windows":
        err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
    default:
        fmt.Println("Unsupported platform. Please open the following URL in your browser:")
        fmt.Println(url)
    }

    if err != nil {
        fmt.Println("Error opening browser:", err)
    }
}
