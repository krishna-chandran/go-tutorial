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


## Initiating a Go Module

1. **Initialize the module:**
    - In your project directory, run:
        ```sh
        go mod init <module-name>
        ```
    - Replace `<module-name>` with the name of your module (e.g., `github.com/username/projectname`).

2. **Add dependencies:**
    - If your project requires external packages, you can add them by running:
        ```sh
        go get <package-name>
        ```
    - Replace `<package-name>` with the name of the package you want to add.

## Running the Go Module

1. **Run the module:**
    - In the terminal, navigate to your project directory and run:
        ```sh
        go run .
        ```

You should see the output of your Go program.



## Running the Go Program with Docker Compose

### Docker Installation

1. **Download Docker:**
    - Visit the [official Docker website](https://www.docker.com/products/docker-desktop).
    - Download the Docker Desktop installer for your operating system.

2. **Install Docker:**
    - Follow the instructions provided by the installer.
    - Verify the installation by opening a terminal and typing:
      ```sh
      docker --version
      ```

### Docker Compose Installation

1. **Download Docker Compose:**
    - Visit the [Docker Compose releases page](https://github.com/docker/compose/releases).
    - Download the appropriate version for your operating system.

2. **Install Docker Compose:**
    - Follow the instructions provided on the releases page.
    - Verify the installation by opening a terminal and typing:
      ```sh
      docker-compose --version
      ```

### Running the Go Program

1. **Create a `Dockerfile`:**
    - In your project directory, create a file named `Dockerfile` and add the following content:
      ```Dockerfile
      FROM golang:1.17-alpine
      WORKDIR /app
      COPY . .
      RUN go build -o main .
      CMD ["./main"]
      ```

2. **Create a `docker-compose.yml` file:**
    - In your project directory, create a file named `docker-compose.yml` and add the following content:
      ```yaml
      version: '3.8'
      services:
        app:
          build: .
          ports:
            - "8080:8080"
      ```

3. **Build and run the application:**
    - In the terminal, navigate to your project directory and run:
      ```sh
      docker-compose up --build
      ```

You should see the output of your Go program in the terminal.


## Getting started with Go Fiber and Postgres

1. **Install the below dependancies
    - In the terminal, navigate to your project directory and run:
      ```
      go get github.com/gofiber/fiber/v2
      go get gorm.io/gorm
      go get gorm.io/driver/postgres
      go get github.com/gofiber/template/html/v2
      ```
