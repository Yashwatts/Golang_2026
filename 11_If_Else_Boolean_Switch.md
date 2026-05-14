# Conditionals (if / else) and Boolean Data Type

Conditionals allow your program to make decisions and execute different blocks of code based on whether a specific condition is true or false.

1. The Boolean Data Type (bool)
Before understanding if statements, you must understand the Boolean type.

• A boolean is a data type that can only have one of two values: true or false.

• Comparison Operators: These operators are used to create boolean expressions:

- == : Equal to (Note: = is for assignment, == is for comparison).
- != : Not equal to.
- > / < : Greater than / Less than.
- >= / <= : Greater than or equal to / Less than or equal to.

• Logical Operators: Used to chain multiple conditions together:

- && (AND): True only if both sides are true.
- || (OR): True if at least one side is true.
- ! (NOT): Reverses the boolean value (Negation).


2. The if Statement Structure
The code inside the if block only runs if the condition evaluates to true.

if remainingTickets == 0 {
    // This code runs only if no tickets are left
    fmt.Println("Our conference is booked out.")
}


3. else and else if

• else: Provides an alternative block of code to run if the if condition is false.

• else if: Allows you to check for multiple specific conditions in a sequence. Only the first true condition in the chain will be executed.


# Key Rules and Logic

Brace Positioning
In Go, the opening curly brace { must be on the same line as the if or else statement. If you place the brace on a new line, the compiler will throw an error.

Conditionals as Variables
You can store the result of a comparison in a boolean variable to make your code more readable:

var noTicketsRemaining bool = remainingTickets == 0
if noTicketsRemaining {
    // Logic here
}

The continue and break keywords
Often used inside conditionals within a loop:

• continue: Skips the rest of the code in the current loop iteration and jumps to the next one (e.g., if user input is invalid).

• break: Exits the loop entirely (e.g., when the conference is sold out).


# Practical Example: Input Validation

For example, checking if a name is long enough AND if an email contains an "@" sign:

isValidName := len(firstName) >= 2 && len(lastName) >= 2
isValidEmail := strings.Contains(email, "@")

if isValidName && isValidEmail {
    // Proceed with booking
} else {
    // Show error message
}


# Switch Statement
The switch statement is a conditional control structure used to execute different blocks of code based on the value of a single variable. It is often used as a cleaner, more readable alternative to a long chain of if-else if statements.

1. Basic Syntax
Instead of checking a variable multiple times with if city == "London", else if city == "Berlin", etc., you "switch" on the variable once.

• switch: The keyword followed by the variable you want to check.
• case: The specific value you are looking for.
• default: A catch-all block that executes if none of the defined cases match the variable's value. This is similar to the final else in an if-else chain.

Example Structure:
switch city {
case "New York":
	// execute code for booking New York conference tickets
case "Singapore", "Hong Kong":
	// execute code for booking Singapore conference tickets
case "London", "Berlin":
	// execute code for booking London conference tickets
case "Mexico":
	// execute code for booking Mexico conference tickets
default:
	fmt.Print("No valid city selected")
}

2. Key Features & Advantages

• Readability: When you have many possible values for a single variable (e.g., 6 different cities), a switch statement is much easier to read and maintain than nested if-else blocks.

• Multiple Values in One Case: You can consolidate multiple values into a single case by separating them with commas. If the variable matches any of those values, that block of code will run.

• No "Fallthrough" by Default: Unlike some languages (like C or Java), Go does not require a break statement at the end of every case. Once a match is found and that case's code is executed, the switch statement finishes automatically.


# When to use Switch vs. If-Else?

• Use if-else: When you are checking different variables or complex ranges (e.g., age > 18 && score < 50).

• Use switch: When you are comparing a single variable against a list of specific, distinct values.