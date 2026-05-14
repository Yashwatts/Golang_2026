# Scope Rules in Go

"Scope" refers to the visibility and lifetime of variables, constants and functions within your code. It determines where a specific name can be accessed and used. Go has three primary levels of scope:

1. **Local (Function) Scope**
   
   Variables defined inside a function are "local" to that function.

   - **Visibility:** They can only be accessed within the curly braces `{}` of the function where they were declared.

   - **Lifetime:** The variable is created when the function is called and destroyed when the function finishes.

   - **Block Scope:** Even within a function, a variable defined inside a specific block (like an if statement or a for loop) is only visible inside that specific block.


2. **Package Scope**
   
   Variables or functions defined outside of any function (at the top level of a file) have package scope.

   - **Visibility:** They are accessible to every function in every file that belongs to the same package (e.g., `package main`).

   - **Requirement:** You cannot use the short-hand `:=` syntax for package-level variables; you must use the `var` or `const` keyword.


3. **Global (Exported) Scope**
   
   Go uses a unique naming convention to handle "global" visibility across different packages.

   - **Exported Names:** If a variable, constant or function name starts with an Uppercase letter (e.g., `ValidateUserInput`), it is exported. This means it can be imported and used by other packages.

   - **Unexported (Private) Names:** If it starts with a lowercase letter (e.g., `userTickets`), it is "unexported" and only visible within its own package.


# Summary of Visibility Rules

| Scope Level | Declaration Location |
|-------------|----------------------|
| Local | Inside a function or block |
| Package | Outside functions (lowercase) |
| Global | Outside functions (Uppercase) |


# Best Practices: "As Local as Possible"

A key software engineering principle: Always define a variable in the smallest scope possible.

- If a variable is only needed inside a loop, define it there.

- If it's needed across a whole function, define it at the top of the function.

- **Why?** This prevents naming conflicts, makes the code easier to debug and ensures memory is cleared as soon as the variable is no longer needed.