# Write our First Program & Structure of a Go File

In Go, every file follows a strict hierarchical structure. If any of these components are missing or out of order, the compiler will throw an error.

1. **The Package Declaration**
   
   The very first line of every Go file must declare which package it belongs to. This is the way Go organizes and reuses code.

   - A package is essentially a container for a collection of source files located in the same directory. All files in that directory must share the same package name.

   - `package main`: This is a special declaration. It tells the Go compiler that this specific file should compile as an executable program rather than a shared library.


2. **Import Statements**
   
   To use code from other packages (like printing text or formatting strings), you must "import" them.

   - `import "fmt"`: The fmt (format) package is part of the Go standard library. It contains functions for formatted Input/Output (I/O).

   - **Strict Rule:** If you import a package but don't use it, Go will throw a compiler error. This keeps binaries small and code clean.

   - **Multiple Imports:** When using more than one package, use a block syntax:
   ```go
   import (
       "fmt"
       "strings"
   )
   ```


3. **The Entry Point: func main()**
   
   Go needs to know exactly where to start execution when you run the program.

   - **Entry Point:** The main function is the mandatory starting point for any executable program.

   - **Constraint:** You can only have one main function in the entire main package.

   - **Syntax Structure:**
   ```go
   func main() {
       // Logic goes here
       fmt.Print("Hello World")
   }
   ```


# Running the Program

You don't always need to build a permanent file to test your code. You can use the "run" command:

```bash
go run main.go
```

- **What happens?** Go compiles the code into a temporary binary, executes it and then deletes the temporary file.


# Key fmt Functions Explained

- `fmt.Print()`: Outputs text to the console. It does not add a new line at the end.

- `fmt.Println()`: "Print Line." Outputs text and automatically adds a new line at the end.

- `fmt.Printf()`: "Print Format." This is more advanced. It allows you to use placeholders (like `%v` for a value or `%T` for a type) to inject variables directly into a string without messy concatenation.


# Explore Packages

To see what each package includes and to find the correct package to import for your project, you should visit the [Go Standard Library Documentation](https://pkg.go.dev/std).

**How to use the Documentation:**

- **Search for Functionality:** If you need to perform a specific task (like string manipulation or time formatting), use the search bar to find the relevant package.

- **Explore Packages:** Clicking on a package, such as the `fmt` package, will show you all available functions (like `Printf` or `Println`), detailed examples of how to use them, and the exact import path needed for your code.

- **Version Info:** The documentation also specifies which version of Go introduced certain features, ensuring compatibility with your local setup.

> **Tip:** Clicking the underlined package names in Visual Studio Code will also take you directly to these official documentation pages.


# Important Conceptual Notes

- **Case Sensitivity:** In Go, if a function name starts with a Capital Letter (like `Println`), it means it is "exported" (public) and can be used by other packages. If it starts with a lowercase letter, it is private to its own package.

- **Brace Style:** Go enforces a specific brace style. The opening curly brace `{` must be on the same line as the function declaration or statement. Putting it on a new line will cause a syntax error.