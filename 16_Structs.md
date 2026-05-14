# Structs

A Struct (short for Structure) is a typed collection of fields. It is used to group related data together to form a single entity. If you are familiar with Object-Oriented Programming (OOP), a struct in Go is very similar to a Class in other languages.

1. **Why use Structs over Maps?**
   
   While maps are useful for storing key-value pairs, they have a major limitation in Go: all values in a map must be of the same data type.

   - **The Problem with Maps:** If you want to store a user's name (string) and their ticket count (int), you have to convert the integer to a string, which is inefficient and messy.

   - **The Struct Solution:** Structs allow you to mix different data types (strings, ints, bools, etc.) within the same structure.


2. **Defining a Struct**
   
   You define the "blueprint" of your data structure outside of your functions.

   - **Syntax:**
   ```go
   type UserData struct {
       firstName       string
       lastName        string
       email           string
       numberOfTickets uint
   }
   ```

   - This creates a custom data type called `UserData`. You are essentially telling Go what a "User" looks like in your specific application.


# Creating and Using Structs

- **Initialization:** You create an "instance" of your struct by assigning values to its fields.

  ```go
  var userData = UserData {
      firstName: firstName,
      lastName: lastName,
      email: email,
      numberOfTickets: userTickets,
  }
  ```

- **Accessing Fields (.):** To get or set a value in a struct, you use the "dot notation."
  
  **Example:** `fmt.Println(userData.firstName)`

- **Type Safety:** Because the struct is a defined type, your IDE (like VS Code) can give you auto-completion suggestions and catch typos in field names before you even run the code.


# Slice of Structs

To keep a list of multiple users where each user is a complex object, you create a slice of your custom struct.

- **Syntax:** `var bookings = make([]UserData, 0)`

- This is much cleaner than a "slice of maps" because it enforces a consistent structure for every entry in your list.


# Key Differences: Map vs. Struct

| Feature | Map | Struct |
|---------|-----|--------|
| Data Types | Only one type for all values | Mixed types allowed |
| Flexibility | Dynamic (add keys anytime) | Fixed (structure is predefined) |
| Performance | Slightly slower (runtime checks) | Faster (compiled checks) |
| Best For | Generic collections / dynamic keys | Entities (User, Product, Event) |