Booleans in Go are represented by the bool type. A bool is either true or false.

Go supports three boolean operators: ! (NOT), && (AND), and || (OR).

true || false // => true
true && false // => false
!true // => false
The three boolean operators each have a different operator precedence. As a consequence, they are evaluated in this order: ! first, && second, and finally ||. If you want to force a different ordering, you can enclose a boolean expression in parentheses (ie. ()), as the parentheses have an even higher operator precedence.

!true && false   // => false
!(true && false) // => true
