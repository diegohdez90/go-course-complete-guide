// documentation for the code in app.go
// This code is a simple Go program that prints "Hello, World!" to the console. It imports the fmt package, which provides functions for formatted I/O operations. The main function is the entry point of the program, and it uses fmt.Println to print the message to the console when the program is executed.
// The code is structured as follows:
// 1. The package main declaration indicates that this is a standalone executable program.
// 2. The import statement imports the fmt package, which is used for formatted I/O operations.
// 3. The main function is defined, which is the entry point of the program. Inside the main function, fmt.Println is called to print "Hello, World!" to the console when the program runs.
// Overall, this code serves as a basic example of a Go program that demonstrates how to use the fmt package to print output to the console.

// In Go, the package name "main" is special because it indicates that the package is an executable program. When you run a Go program, the Go runtime looks for a package named "main" and executes the main function within that package. If the package is not named "main", it cannot be executed directly and is typically used as a library or a module that can be imported by other packages. Therefore, naming the package "main" is essential for creating an executable Go program.
// In this code, the package is named "main" because it contains the main function, which is the entry point of the program. This allows the program to be executed directly, and when run, it will print "Hello, World!" to the console.
// In summary, the package is called "main" because it signifies that this package is an executable program, and it contains the main function that serves as the entry point for the program's execution.
package main

//
// In Go, fmt is a package that is imported and used in the code. It is not a variable but a package name.
// The fmt package provides functions for formatted I/O operations, such as printing to the console.
// In this code, fmt.Println is used to print "Hello, World!" to the console.
import "fmt"

// explain why is set as main function
//
// In Go, the main function is the entry point of the program. When you run a Go program, the execution starts from the main function.
// The main function is defined in the main package, and it is where you can write the code that you want to execute when the program runs.
// In this code, the main function is used to print "Hello, World!" to the console when the program is executed.
func main() {
	fmt.Println("Hello, World!")
}

// steps to run a build a go program from creating a module to running the program
//
// 1. Create a new directory for your Go project and navigate to it in the terminal.
// 2. Initialize a new Go module by running the command: go mod init <module-name>. This will create a go.mod file in your project directory.
// 3. Create a new Go file (e.g., main.go) and write your Go code in it.
// 4. To build the program, run the command: go build. This will compile your code and create an executable file in the same directory.
// 5. To run the program, use the command: ./<executable-name> (on Unix-based systems) or <executable-name>.exe (on Windows). This will execute your program and you should see the output in the console.
