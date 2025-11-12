/*
Go is a statically typed, compiled programming language. Go is syntactically similar to C.
The language is often referred to as Golang, but the proper name is Go.

The Basics concept introduces three major language features: Packages, Functions, and Variables.

--- Packages ---
Go applications are organized in packages. A package is a collection of source files in the same directory.
- All source files in a directory must share the same package name.
- Only entities (functions, types, variables, constants) starting with a capital letter (PascalCase) are accessible across packages.
- Identifiers are conventionally named using camelCase, except for accessible entities, which use PascalCase.

--- Variables ---
Go is statically-typed, so all variables must have a defined type at compile-time.
- Once declared, a variable's type can never change.

--- Constants ---
Constants hold data whose value cannot change during program execution.
- Defined using the 'const' keyword.

--- Functions ---
Go functions accept zero or more parameters. Parameters must be explicitly typed (no type inference).
- Values are returned from functions using the 'return' keyword.

--- Comments ---
Go supports two types of comments: Single line (//) and multiline (/* */).
*/
package lasagna

import "fmt"

// --- Constants ---
const OvenTime = 40 // Defines a numeric constant 'OvenTime' with the value of 40
const Age = 21 // Defines a numeric constant 'Age' with the value of 21

// --- Variables ---

// VariableDeclaration demonstrates explicit and implicit declaration.
func VariableDeclaration() {
	// Explicitly typed: var explicit int // Explicitly specifying a type
	var explicit int = 5 

	// Implicitly typed: The compiler assigns the variable type to match the initializer.
	implicit := 10   // Implicitly typed as an int

	count := 1 // Assign initial value
	count = 2  // Update to new value (must match the original int type)

	// count = false // This throws a compiler error due to assigning a non `int` type
	fmt.Printf("Explicit: %d, Implicit: %d, Count: %d\n", explicit, implicit, count)
}

// ExpectedMinutesInOven is a public function.
// Functions accept zero or more parameters.
func ExpectedMinutesInOven() int {
	// Values are returned from functions using the return keyword.
	return OvenTime 
}

// hi is a private function.
// Parameters must be explicitly typed (name string).
func hi(name string) string {
	return "hi " + name
}

// Hello is a public function demonstrating function invocation.
func Hello(name string) string {
	// A function is invoked by specifying the function name and passing arguments.
	return hi(name) 
}
