# Local Setup - Install Go & Editor

To start developing in Go, you need two main software components: the Go Compiler and an Integrated Development Environment (IDE).

1. **The IDE: Visual Studio Code**
   
   While you can use any text editor, the tutorial recommends Visual Studio Code (VS Code) because it is free, lightweight and has excellent community support for Go.

   - **Download:** Visit the official VS Code website (https://code.visualstudio.com) and download the installer for your operating system (Windows, macOS, or Linux).

   - **Go Extension:** After installing VS Code, you must install the official Go extension (developed by the Google Go team).
     - Open the Extensions tab (square icon on the left sidebar).
     - Search for "Go" and click Install.
     - Why? It provides syntax highlighting, intellisense (auto-completion) and code navigation tools.

2. **The Go Compiler**
   
   The compiler is what translates your human-readable code into machine code that the computer can run.

   **Installation:**
   1. Go to go.dev/dl to download the package for your OS.
   2. Run the installer and follow the default prompts.

   **Verification:**
   1. Open your terminal (Command Prompt, PowerShell or macOS Terminal).
   2. Type `go version` and hit enter.
   3. If correctly installed, the terminal will display the installed version (e.g., `go version go1.26.3 windows/amd64`).


# Project Initialization

Once your tools are ready, you need to turn your working folder into a Go Module.

1. **Create a folder:** In your terminal, create a new directory (e.g., `mkdir booking-app`).

2. **Open in VS Code:** Use `code .` in the terminal to open the folder.

3. **Initialize Module:** Open the integrated terminal in VS Code and run: `go mod init <project-name>`
   - Example: `go mod init booking-app`
   - This creates a `go.mod` file, which tracks your project's dependencies and version information.


# Pro-Tip: "Install All"

When you first open a `.go` file in VS Code, a popup will likely appear in the bottom right corner asking to install additional tools (like `gopls` or `dlv`). Always click "Install All" to ensure your development environment has full debugging and formatting capabilities.