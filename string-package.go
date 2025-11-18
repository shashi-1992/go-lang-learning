The strings package contains many useful functions to work on strings.

import "strings"

strings.ToUpper("test") // => "TEST"
=============================================================
About Strings Package
The strings package contains many useful functions to work on strings:

Function	Purpose
ToLower	Convert the string to lower case
ToUpper	Convert the string to upper case
TrimSpace	Trim leading and trailing whitespace
Index	Find the index of the first instance of a substring within a string
Replace	Replace one occurrence of a substring in a string
ReplaceAll	Replace all occurrences of a substring in a string
Split	Split a string into parts based on a separator
HasSuffix	Check if a string ends with a specific substring
Count	Count the number of occurrences of a substring within a string
import "strings"

strings.ToLower("Gopher")    // => "gopher"
strings.Index("Apple", "le") // => 3
strings.Count("test", "t")   // => 2
