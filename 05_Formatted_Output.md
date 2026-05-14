# Formatted Output - printf

The fmt.Printf() function (Print Format) provides more control over how text and variables are displayed in the console compared to fmt.Println(). It allows you to build complex strings using annotation markers or placeholders.

1. Why use printf?
When using fmt.Println(), you often have to separate strings and variables with commas, which automatically adds spaces and can be messy for complex sentences:

• Println style: fmt.Println("Welcome to", conferenceName, "booking application")

• Printf style: fmt.Printf("Welcome to %v booking application\n", conferenceName)

2. Placeholders (Verbs)
Placeholders tell Go where to inject a variable's value and how to format it.

• %v (Value): The default format. It prints the value of the variable in its natural format.

• %T (Type): Prints the data type of the variable (e.g., string, int, uint).

• \n (New Line): Unlike Println, Printf does not add a new line at the end automatically. you must add \n manually at the end of your string to move the cursor to the next line.

3. Usage & Order
When using multiple placeholders, the variables must be passed in the exact order they appear in the string.

• Example:
fmt.Printf("We have total of %v tickets and %v are still available\n", totalTickets, remainingTickets)

• The first %v is replaced by totalTickets.
• The second %v is replaced by remainingTickets.


# Advanced Formatting & Documentation
Go offers a vast library of specific placeholders for different data types (e.g., specific precision for floats or binary representations of integers).

• Where to find them: You can view the full list of "verbs" in the official fmt package documentation (https://pkg.go.dev/fmt).

• Hover Feature: In VS Code, you can hover over the fmt import or the Printf function to see a direct link to the documentation.