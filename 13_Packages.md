# Organizing Code with Go Packages

As your application grows in complexity, keeping all your code in a single main.go file becomes unmanageable. Go uses Packages to logically group related code, making it easier to maintain and share across projects.

1. What is a Package?
A package is essentially a directory containing one or more Go source files.

• All files in the same folder must belong to the same package.

• The main Package: This is the entry point of your application.

• Standard Library Packages: Go comes with built-in packages like fmt (formatting), strings (string manipulation), and math.


2. Creating Your Own Packages
To create a custom package:

• Create a new folder: The folder name typically matches the package name (e.g., a folder named helper for a helper package).

• Move logic: Move related functions (like validation logic) into a new .go file within that folder.

• Declare Package: The first line of the new file must be package <name> (e.g., package helper).


# Exporting and Importing

Exporting (Public vs. Private)
Go uses a very simple rule for visibility:

• Capitalized Name: If a function, variable, or struct starts with a Capital Letter, it is "Exported" (public) and can be accessed from other packages.

• Lowercase Name: If it starts with a lowercase letter, it is "Unexported" (private) and can only be used within its own package.

• Example: func ValidateInput is public; func validateInput is private.


Importing Your Own Package
To use your custom package in your main package, you must import it using its full path:

• Path Structure: <module-name>/<package-name>

• Example: If your module is named booking-app and your package is in the helper folder, you import it as:
import "booking-app/helper"

• Usage: Call the exported function using the package name prefix: helper.ValidateInput().


# Important Scoping Rules

| Scope Level | Description |
| Local / Function Scope | Variables defined inside a function are only accessible within that function. |
| Package Scope | Variables/Functions defined outside functions but starting with a lowercase letter are accessible to all files in the same package. |
| Global / Exported Scope | Starting a variable/function with a capital letter makes it accessible to any package that imports it. |


# Why Use Multiple Packages?

• Logical Grouping: Grouping "Helper" functions in one place and "Database" logic in another makes the project structure clear.

• Namespace Management: It prevents naming conflicts. You can have a Validate() function in a user package and a Validate() function in a ticket package without them clashing.

• Reusability: You can easily copy a well-written package folder into a completely different project.