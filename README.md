# go-tutorial


# Go Tutorial

## Installation Instructions

1. **Download Go:**
    - Visit the [official Go website](https://golang.org/dl/).
    - Download the installer for your operating system.

2. **Install Go:**
    - Follow the instructions provided by the installer.
    - Verify the installation by opening a terminal and typing:
      ```sh
      go version
      ```

## Running a Go Program

1. **Create a Go file:**
    - Create a new file named `main.go` in your project directory:
      ```sh
      touch main.go
      ```

2. **Write a simple Go program:**
    - Open `main.go` in a text editor and add the following code:
      ```go
      package main

      import "fmt"

      func main() {
            fmt.Println("Hello, World!")
      }
      ```

3. **Run the Go program:**
    - In the terminal, navigate to your project directory and run:
      ```sh
      go run main.go
      ```

You should see the output:
```
Hello, World!
```

## Platform-Specific Installation Instructions

### Windows

1. **Download Go:**
    - Visit the [official Go website](https://golang.org/dl/).
    - Download the Windows installer.

2. **Install Go:**
    - Run the downloaded installer and follow the prompts.
    - Verify the installation by opening Command Prompt and typing:
      ```sh
      go version
      ```

### Linux

1. **Download Go:**
    - Visit the [official Go website](https://golang.org/dl/).
    - Download the Linux tarball.

2. **Install Go:**
    - Extract the tarball to `/usr/local`:
      ```sh
      sudo tar -C /usr/local -xzf go1.xx.x.linux-amd64.tar.gz
      ```
    - Add Go to your PATH by adding the following line to your `.profile` or `.bashrc`:
      ```sh
      export PATH=$PATH:/usr/local/go/bin
      ```
    - Apply the changes:
      ```sh
      source ~/.profile
      ```
    - Verify the installation by opening a terminal and typing:
      ```sh
      go version
      ```

### macOS

1. **Download Go:**
    - Visit the [official Go website](https://golang.org/dl/).
    - Download the macOS installer.

2. **Install Go:**
    - Run the downloaded installer and follow the prompts.
    - Verify the installation by opening Terminal and typing:
      ```sh
      go version
      ```
