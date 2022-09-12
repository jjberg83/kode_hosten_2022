package errors

import(
	//"fmt"
	"strconv"
)

/*
Task 5: Errors needed for multiwriter

You may find this blog post useful:
http://blog.golang.org/error-handling-and-go

Similar to a the Stringer interface, the error interface also defines a
method that returns a string.

type error interface {
    Error() string
}

Thus also the error type can describe itself as a string. The fmt package (and
many others) use this Error() method to print errors.

Implement the Error() method for the Errors type defined above.

The following conditions should be covered:

1. When there are no errors in the slice, it should return:

"(0 errors)"

2. When there is one error in the slice, it should return:

The error string return by the corresponding Error() method.

3. When there are two errors in the slice, it should return:

The first error + " (and 1 other error)"

4. When there are X>1 errors in the slice, it should return:

The first error + " (and X other errors)"
*/
func (m Errors) Error() string {
	// lag loop som går gjennom Errors, en slice med Errors
	antall_errors := 0
	kun_feil_slice := []error{}
	for _,err := range m { 
		if err != nil {
			antall_errors++
			kun_feil_slice = append(kun_feil_slice, err)
		}
	}
	if antall_errors == 0 {
		return "(0 errors)"
	} else if antall_errors == 1 {
		return kun_feil_slice[0].Error()
	} else if antall_errors == 2 {
		return kun_feil_slice[0].Error()  + " (and 1 other error)"
	} else {
		return kun_feil_slice[0].Error() + " (and " + strconv.Itoa(len(kun_feil_slice) - 1) +" other errors)"
	}
}
