Booleans in Go are represented by the bool type. A bool is either true or false.

Go supports three boolean operators: ! (NOT), && (AND), and || (OR).

true || false // => true
true && false // => false
!true // => false
The three boolean operators each have a different operator precedence. As a consequence, they are evaluated in this order: ! first, && second, and finally ||. If you want to force a different ordering, you can enclose a boolean expression in parentheses (ie. ()), as the parentheses have an even higher operator precedence.

!true && false   // => false
!(true && false) // => true

===============================
Booleans in Go are represented by the predeclared boolean type bool, which values can be either true or false. It's a defined type.

var closed bool    // boolean variable 'closed' implicitly initialized with 'false'
speeding := true   // boolean variable 'speeding' initialized with 'true'
hasError := false  // boolean variable 'hasError' initialized with 'false' 
Go supports three logical operators that can evaluate expressions down to Boolean values, returning either true or false.

Operator	What it means
&& (and)	It is true if both statements are true.
|| (or)	It is true if at least one statement is true.
! (not)	It is true only if the statement is false.
