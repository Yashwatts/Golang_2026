# Encapsulate Logic with Functions

As an application grows, the main function can become crowded and difficult to read. Functions allow you to group related code into a named container, making the code reusable, maintainable and organized.

1. **Function Declaration and Syntax**
   
   To create a function, use the `func` keyword followed by a descriptive name and a block of code.

   - **Syntax:**
   ```go
   func functionName() {
       // Logic goes here
   }
   ```

   - **Note:** Simply defining a function doesn't run it. You must explicitly call it inside the main function (or another function) using its name: `greetUsers()`.


2. **Input Parameters**
   
   Functions often need data from the "outside" to do their job. You pass this data through parameters.

   - **Defining Parameters:** You must specify the name and the data type for every parameter.

   - **Example:**
   ```go
   func greetUsers(confName string, confTickets int) {
       fmt.Printf("Welcome to %v!", confName)
   }
   ```

   - This makes the function flexible, as it can behave differently based on the data you hand over to it.


3. **Returning Values (Output)**
   
   A function can perform a calculation and "return" the result back to the place it was called.

   - **Return Type:** You must declare the data type of the value being returned after the parentheses.

   ```go
   func getFirstNames(bookings []string) []string {
       // Logic to extract names
       return firstNames
   }
   ```

   - **Multiple Return Values:** Uniquely, Go allows functions to return multiple values at once. This is commonly used to return both a result and an error status.

   **Syntax:**
   ```go
   func validateInput(...) (bool, bool, bool) { ... }
   ```


# Key Concepts & Best Practices

**Local vs. Package Level Variables**

- **Local Variables:** Defined inside a function. They are only "visible" and usable within that specific function.

- **Package Level Variables:** Defined outside of any function. They are accessible to all functions within that same package, which reduces the need to pass many parameters back and forth.


**Code Reuse**

The primary power of functions is reuse. If you have a complex validation logic, you write it once in a function and call it every time a user submits a form, rather than rewriting the same if-else blocks repeatedly.

**Refactoring**

The process of taking a large block of code and breaking it into smaller, descriptive functions is called "Refactoring." It makes the main function look like a clean list of high-level steps rather than a mess of technical details.


# Example Refactored Main Function

After refactoring, your main function should look clear and readable:

```go
func main() {
    greetUsers()
    firstName, lastName, email, userTickets := getUserInput()
    isValid := validateUserInput(firstName, lastName, email, userTickets)

    if isValid {
        bookTicket(userTickets, firstName, lastName, email)
        printFirstNames()
    }
}
```