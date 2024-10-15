/**
 * @file main.go
 * @brief Entry point for the Go application.
 *
 * The `package main` is required in Go to define an executable program.
 * The `main` package is a special package in Go that serves as the entry point
 * for the application. When you run a Go program, the `main` function within
 * the `main` package is executed.
 *
 * Naming the package `main` is necessary for creating an executable file.
 * If you name the package anything other than `main`, the Go compiler will
 * generate a package archive (a `.a` file) instead of an executable.
 */
package main //

import "fmt" 
/**
 * @brief Importing necessary packages.
 *
 * The `import` keyword is used to include the necessary packages in the Go program.
 * Packages are used to organize and reuse code. The `fmt` package is imported here
 * to provide formatted I/O functions, such as printing to the console.
 */

func main() {
	fmt.Println("Hello, World!")
}