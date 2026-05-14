# Maps

A Map is a built-in data type in Go that stores a collection of key-value pairs. It is similar to a dictionary in Python, an object in JavaScript, or a HashMap in Java.

1. Why use Maps?
the bookings list only stored strings (the full names). However, a real booking needs more data: first name, last name, email and the number of tickets.

• The Limitation of Slices: A slice can only store a list of a single type (e.g., all strings).

• The Map Solution: A map allows you to group different pieces of information together under unique "keys" (e.g., Key: "email", Value: "user@example.com").


2. Creating and Initializing a Map
You must define the data type for both the Key and the Value.

• Syntax: map[KeyType]ValueType

• Using make(): In Go, you cannot simply declare a map variable; you must initialize it using the make() function to allocate memory.

• Example:
var userData = make(map[string]string)
This creates an empty map where both the keys and values are strings.


# Working with Map Data

• Adding/Updating Data: Use square brackets with the key name.

userData["firstName"] = firstName
userData["lastName"] = lastName
userData["email"] = email

• The Type Limitation: A single map cannot have mixed data types for its values. If you define a map[string]string, every value must be a string.

• The Workaround: To store an integer (like userTickets) in a string-only map, you must convert the integer to a string first using the strconv package.


# Slice of Maps

To store a list of multiple users where each user has their own set of key-value pairs, you create a Slice of Maps.

• Syntax: var bookings = make([]map[string]string, 0)

• Adding to the list: You use the append() function to add a user's map to the bookings slice.


# Key Differences: Map vs. Array/Slice

| Feature | Array / Slice | Map |
| Access | Via numeric Index (0, 1, 2...) | Via unique Key ("email", "id"...) |
| Order | Elements are ordered | Elements are unordered |
| Usage | Lists of identical items | Complex entities with properties |


# Important Note on Go Map Restrictions

Because Go is statically typed, the restriction of having only one value type in a map (e.g., all strings) can be frustrating when you want to mix strings, integers and booleans. To solve this properly without messy type conversions, Go provides Structs.