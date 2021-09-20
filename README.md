# go

## Types
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

## Constants
Constants can be character, string, boolean, or numeric values.

## Slices
Slices are like references to arrays
A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.

[John Paul George Ringo]
[John Paul] [Paul George]
[John XXX] [XXX George]
[John XXX George Ringo]

Slice length and capacity
A slice has both a length and a capacity.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

## Pointer receivers
Methods with pointer receivers can modify the value to which the receiver points. 
Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

With a value receiver, the method operates on a copy of the original value.

## Loops
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }


