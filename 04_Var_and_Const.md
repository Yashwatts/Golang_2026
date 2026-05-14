# Variables & Constants in Go

1. **Variables (var)**
   
   Variables are used to store values that can change during the execution of a program. Instead of repeating the same value throughout your code, you store it once in a variable and reference its name.

   - **Syntax:** `var variableName type = value`

   - **Example:** `var conferenceName = "Go Conference"`

   - **Naming Convention:** Go uses camelCase (e.g., `remainingTickets`). It is best practice to use descriptive names that explain what the value represents.

   - **Reassignment:** You can update a variable's value later in the code:
   ```go
   remainingTickets = remainingTickets - userTickets
   ```


2. **Constants (const)**
   
   Constants are used for values that cannot be changed once they are assigned. This provides a safety layer to ensure critical values (like the total number of tickets) aren't accidentally overwritten.

   - **Syntax:** `const constantName = value`

   - **Example:** `const totalTickets = 50`

   - **Rule:** If you try to reassign a value to a constant, the Go compiler will throw an error: "Cannot assign to [constant name]".


# The "Syntactic Sugar" (Short Hand)

Go provides a more concise way to declare and initialize variables without using the `var` keyword or explicitly stating the type.

- **Short Declaration:** `variableName := value`

- **Example:** `conferenceName := "Go Conference"`

- **Limitations:**
  - It cannot be used for constants.
  - It cannot be used if you want to declare a variable without assigning a value immediately (e.g., when waiting for user input).
  - It cannot be used for package-level variables (outside of functions).


# Key Go-Specific Behaviors

**Variable Usage Requirement**

In Go, if you declare a variable but do not use it anywhere in your code, the program will not compile. This is a built-in feature to keep code clean and prevent "dead code" or wasted memory.

**Implicit vs. Explicit Typing**

- **Implicit (Inferred):** If you assign a value immediately, Go automatically "infers" the type. (e.g., `var name = "Yash"` is automatically a string).

- **Explicit:** If you don't provide a value immediately, you must define the type:
```go
var userName string
var userTickets int
```

**The Concept of Pointers (&)**

When getting user input, you don't pass the variable itself; you pass its memory address using a pointer.

- **Syntax:** `&variableName`

- **Purpose:** This tells Go where in the computer's memory the variable is stored so that a function (like `fmt.Scan`) can write the user's input directly into that specific spot.