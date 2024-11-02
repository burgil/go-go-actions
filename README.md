# go-go-actions

Go, go Action! Testing GitHub actions build capabilities because code signing certificates are too expensive

```bash
go mod init localhost/m/v2
```

**The goal of this project is to test if building through GitHub Actions can bring benefits to my development environment.**

## About the (example) project

This is a simple go web application that sets up a web server and opens the default browser based on your platform.

## what this app does

- starts a web server on port 8080.
- serves a simple "hello, world!" message at the root URL.
- automatically opens the default web browser to the server's address when the application runs.

## how to run the project

1. make sure you have go installed on your machine.
2. clone the repository or download the project files.
3. navigate to the project directory in your terminal.
4. run the following command:

   ```bash
   go run main.go
   ```

5. your default browser should open to `http://localhost:8080`.

## platforms supported

The app detects your operating system and opens the appropriate command to launch the browser:

- Linux
- Windows
- MacOS

if your platform is unsupported, the app will provide the URL to open manually.

## dependencies

This project uses the standard library in go, so no additional dependencies are required.

## license

This project is open source and available under the [MIT license](LICENSE).
