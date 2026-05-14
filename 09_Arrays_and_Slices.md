# Arrays & Slices

In Go, arrays and slices are used to store a list of elements of the same data type.

1. Arrays
Arrays have a fixed size that must be defined at the time of creation. This size cannot change while the program is running.

• Syntax: var variableName [size]type

• Example: var bookings [50]string

• Key Characteristics:

- Fixed Length: If you define an array of 50, it will always take up space for 50 elements, even if you only use one.
- Zero-Based Indexing: The first element is at index 0, the second at 1, and so on.
- Direct Assignment: You can assign a value to a specific spot using the index: bookings[0] = "Yash".


2. Slices
A slice is a more flexible, dynamic version of an array. It is an abstraction built on top of arrays that allows the list to grow or shrink automatically. In real-world Go development, slices are used much more frequently than arrays.

• Syntax: var variableName []type (Note the empty square brackets).

• Initialization: Unlike arrays, you don't give a slice a fixed size. You can initialize an empty slice like this: bookings := []string{}.

• Adding Elements (append): To add a value to a slice, you use the built-in append() function. It takes the original slice and the new value, then returns a new, updated slice.

Example: bookings = append(bookings, firstName + " " + lastName)


# Differences at a Glance

| Feature | Array | Slice |
| Size | Fixed (must be defined) | Dynamic (grows as needed) |
| Efficiency | Manual management of indices | More efficient and easier to use |
| Syntax | [50]string | []string |
| Adding Data | bookings[0] = "Value" | append(bookings, "Value") |


# Useful Built-in Functions

• len() (Length): Returns the number of elements currently in the array or slice.

- For Arrays: Returns the fixed size (e.g., 50).
- For Slices: Returns the number of items actually stored.


# Important Conceptual Note
Go does not allow mixed data types in these lists. If you create a []string, you cannot store an integer inside it. This strictness ensures that when you iterate through the list later, you know exactly what type of data to expect.