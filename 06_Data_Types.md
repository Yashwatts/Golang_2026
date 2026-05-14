# Data Types in Go
In Go, every value has a specific Data Type. This tells the compiler how much memory to allocate and what operations (like math or string manipulation) can be performed on that value.

1. Basic Data Types

• Strings (string): Used for textual data. Values must be enclosed in double quotes (e.g., "Go Conference").

• Integers (int): Used for whole numbers, both positive and negative (e.g., -5, 0, 50).

• Booleans (bool): Represents logical values: either true or false.

2. Numeric Type Variations
Go provides multiple versions of numeric types to allow for precise memory management and data validation.

| Type | Description |
| int | The standard integer type. The size (32 or 64 bits) depends on the platform. |
| uint | Unsigned Integer. It only stores positive whole numbers (0 and up). Ideal for values that can never be negative, like ticket counts. |
| float32 / float64 | Used for numbers with decimals (e.g., 3.14). float64 provides higher precision and is the default for floating-point numbers. |


# Key Theoretical Concepts

Type Inference
If you declare a variable and assign a value immediately, Go automatically detects the type.

• var name = "Yash" -> Go knows this is a string.
• var count = 50 -> Go knows this is an int.

Explicit Type Definition
If you do not assign a value immediately (e.g., when waiting for user input), you must specify the type. This prevents bugs by ensuring the variable only ever holds the correct type of data.

• Syntax: var userName string

Type Safety & Conversion
Go is a strongly typed language. You cannot perform operations between different types (e.g., adding an int to a uint) without explicitly converting them first.

• The Problem: Mixing types causes a "type mismatch" error.

• The Solution: Use conversion functions or ensure all variables in a calculation share the same type (e.g., making both remainingTickets and userTickets the type uint).


# Pro-Tip: Checking Types with %T

If you are ever unsure what type Go has assigned to a variable, you can use the %T placeholder in a printf statement to inspect it:

fmt.Printf("conferenceName is type %T\n", conferenceName) 
// Output: conferenceName is type string