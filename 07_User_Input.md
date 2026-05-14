# Getting User Input
To make an application interactive, you need to accept data from the user during runtime. Go provides built-in functions in the fmt package to handle this.

1. The fmt.Scan Function
The Scan function is used to read text entered into the console (standard input).

• Behavior: It scans the input for the next space-separated value.

• Waiting for Input: When Scan is called, the program's execution stops and waits for the user to type something and press Enter.


2. Passing by Reference (Pointers)
A critical detail when using Scan is that you must pass the memory address of the variable where the input should be stored.

• Syntax: Use the ampersand (&) before the variable name.

• Example: fmt.Scan(&firstName)

• If you pass just firstName, you are passing a copy of the current value (which is likely empty). By using &, you give the function a Pointer to the actual spot in memory, allowing it to update the variable directly.


3. Step-by-Step Implementation

• Prompt the User: Use fmt.Print or fmt.Println to tell the user what to type.

• Scan the Input: Use fmt.Scan with a pointer to store the result.
var firstName string
fmt.Println("Enter your first name: ")
fmt.Scan(&firstName)


# Limitations of fmt.Scan

• Space Delimitation: fmt.Scan stops reading at the first whitespace. If a user enters "Virat Kohli" into a single Scan call meant for a name, it will only capture "Virat". The "Kohli" will remain in the input buffer and potentially cause issues for the next Scan call.

• Data Types: If you try to scan a string into an int variable, it may cause an error or unexpected behavior.


# Key Concepts to Remember

• Imports: Ensure import "fmt" is at the top of your file.

• Variables: You must declare the variable (with a type) before you can scan a value into it.

• Order of Execution: The program executes line-by-line. If your scan comes before your prompt, the user will see a blank screen and won't know they need to type anything.