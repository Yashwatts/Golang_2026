# Loops in Go

Loops are used to execute a block of code multiple times. In Go, the approach to loops is simplified: there is only one loop keyword, which is for. Unlike other languages, Go does not have while or do-while loops; the for loop is flexible enough to handle all those use cases.

1. The Infinite Loop
To keep an application running indefinitely (e.g., a booking system that keeps asking for the next user), you can use a for loop without any conditions.

• Syntax:
for {
    // This code repeats forever until the program is stopped
}

• This is equivalent to for true {} in other languages. It creates a continuous loop that only ends if you manually interrupt it (Ctrl+C) or use a break statement.


2. Iterating Over Collections (range)
When you want to go through a slice or an array, you use the for range syntax.

• Syntax: for index, element := range bookings { ... }

• The Return Values: range returns two values for every iteration:
1. Index: The current position in the list (0, 1, 2...).
2. Element: The actual data stored at that position.

• The Blank Identifier (_): In Go, you cannot have unused variables. If you only need the element and not the index, you must replace the index variable with an underscore.

Example: for _, booking := range bookings { ... }


# Advanced Loop Concepts

The "Fields" Function for Strings
A loop to extract first names from a list of full names. To do this, it introduces
strings.Fields():

• This function splits a string at every whitespace and returns a slice of the parts.

• Example: "Virat Kohli" becomes ["Virat", "Kohli"]. By taking index 0, you get the first name.


Conditional Loops
You can also write a loop that only runs while a specific condition is true.

• Synatx:
for remainingTickets > 0 {
    // Runs as long as there are tickets left
}


# Loop Control Statements

• break: Completely stops the loop and exits. For example, if all tickets are sold out, you "break" to end the booking process.

• continue: Skips the current iteration and jumps straight to the next one. This is useful if a user enters invalid data; you skip the booking logic and ask for their name again.


# Summary Table

| Loop Type | Usage | Syntax Example |
| Infinite | Continuous execution | for { ... } |
| Conditional | Run while true | for tickets > 0 { ... }
| Iterative | Loop through a list | for index, val := range list { ... } |