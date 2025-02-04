# Section 08 - Lab 01 : Playing with pointers

Compare array bytes vs slice of byte in pointer to struct.

## TODO - Update Document.Data to use []byte instead of [100e6]byte

Update the example application form Sec08 Lec04 Example 06, such that the Document structure using a []byte instead of an array of []. The Document.Data must still be initialized with an 100 MegaByte slice.

Compare the runtimes between the two implementations. Is there a significant performance differnce between the pointer to Document.Data using an array bytes vs a pointer to Document.Data using []byte?

## Solution

An array with pointers is slightly faster than a slice with pointers, possibly because of a slight overhead from the slice header, or since slices are allocated on heap (vs stack for arrays).

However, there isn't a clear advantage of using slices with vs without pointers because by default only a shallow copy or a slice (vs a deep copy for arrays) is passed through functions.

Array (no pointers)
```
% go run array-no-pointers/main.go
Size of aDoc object: 100000112
Size of object passed to countWords(): 100000112
Word Count: 14937
Size of object passed to capitalizeHeading(): 100000112
Size of object passed to listHeadings(): 100000112
Runtime: 307.169949ms
```
Array (pointers)
```
% go run array-pointers/main.go 
Size of aDoc object: 8
Size of object passed to countWords(): 8
Word Count: 69817
Size of object passed to capitalizeHeading(): 8
Size of object passed to listHeadings(): 8
Runtime: 40.736µs
```
Slice (no pointers)
```
%  go run slice-no-pointers/main.go
Size of aDoc object: 136
Size of object passed to countWords(): 136
Word Count: 114191
Size of object passed to capitalizeHeading(): 136
Size of object passed to listHeadings(): 136
Runtime: 154.245µs
```
Slice (pointers)
```
% go run slice-pointers/main.go
Size of aDoc object: 8
Size of object passed to countWords(): 8
Word Count: 137514
Size of object passed to capitalizeHeading(): 8
Size of object passed to listHeadings(): 8
Runtime: 131.307µs
```
