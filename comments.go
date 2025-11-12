package comments

import "fmt"

// Go uses comments to explain code and make it easier to maintain.
// There are two types of comments in Go: single-line and multi-line.

// SingleLineComment demonstrates how to use // for comments.
// Single-line comments are the most common type and start with two forward slashes (//).
func SingleLineComment(name string) {
	// A single-line comment extends to the end of the line.
	// You should use them to explain the purpose of variables, functions, or complex logic.
	fmt.Printf("Hello, %s!", name) // This print statement greets the user.
}

/*
MultiLineComment demonstrates the use of /* */ for block comments.
This style is useful for:
1. Explaining lengthy logic or algorithms.
2. Temporarily commenting out large blocks of code.
3. Generating documentation (though specific Go Doc comments are preferred).
*/
func MultiLineComment(age int) string {
	// The return value will be the user's age formatted as a string.
	// We use the short variable declaration syntax (:=) here.
	result := fmt.Sprintf("User is %d years old.", age)

	// Note: Block comments are less frequently used than single-line comments in Go source code.
	return result
}
