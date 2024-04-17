package main

import (
	"fmt"
	"slices"
)

// This is the main entry of the application.
func main() {
	// Array Declarations
	// ------------------
	fmt.Println("Array Declaration:")

	// Basic array declaration
	var arrA [3]int
	// Array declaration with Array literal
	arrB := [3]int{10, 20, 30}
	// Sparse array with specific non-zero elements
	arrC := [12]int{1, 5: 4, 6, 10, 100, 15}
	// Array declaration with ... as length
	arrD := [...]int{100, 200, 300}
	// Array declaration with :=
	arrE := [...]int{250, 350, 450, 1000}

	fmt.Println("arrA =", arrA)
	fmt.Println("arrB =", arrB)
	fmt.Println("arrC =", arrC)
	fmt.Println("arrD =", arrD)
	fmt.Println("arrE =", arrE)
	fmt.Println()

	// Comparing Arrays
	// ----------------
	fmt.Println("Comparing Arrays:")

	arrF := [...]int{1, 2, 3, 4, 5}
	arrG := [5]int{1, 2, 3, 4, 5}

	fmt.Println("arrF == arrG:", arrF == arrG)
	fmt.Println()

	// Multidimensional Array
	// ----------------------
	fmt.Println("Multidimensional Arrays:")

	// Array of array of integers: 5 x 5
	var arrH [5][5]int

	fmt.Println(arrH)
	fmt.Println()

	// Array Access
	// ------------
	fmt.Println("Array Access:")

	arrI := [3]int{10, 20, 30}

	// Accessing the values
	fmt.Println("arrI[0] =", arrI[0])
	fmt.Println("arrI[1] =", arrI[1])
	fmt.Println("arrI[2] =", arrI[2])

	// Changing the values
	arrI[0] = 15
	arrI[1] = 25
	arrI[2] = 35

	fmt.Println("arrI =", arrI)
	fmt.Println()

	// Length of an array
	// ------------------
	fmt.Println("Length of an array:")
	arrJ := [3]int{10, 20, 30}
	arrK := [...]int{15, 25, 35}

	fmt.Println("arrJ =", arrJ, "length =", len(arrJ))
	fmt.Println("arrK =", arrK, "length =", len(arrK))
	fmt.Println()

	// Slice Declarations
	// ------------------
	fmt.Println("Slice Declarations:")

	// Slice declaration with Slice literal
	sliceA := []int{10, 20, 30, 40}
	// Slice declaration with specific non-zero elements
	sliceB := []int{1, 5: 4, 6, 10, 100, 15}
	// The Zero-Value of a slice is nil
	var sliceNil []int

	fmt.Println("sliceA =", sliceA)
	fmt.Println("sliceB =", sliceB)
	fmt.Println("sliceNil =", sliceNil)
	fmt.Println("sliceNil is nil?", sliceNil == nil)
	fmt.Println()

	// Multidimensional Slice
	// ----------------------
	fmt.Println("Multidimensional Slices:")

	// Slice of slice of integers
	var sliceC [][]int

	fmt.Println(sliceC)
	fmt.Println()

	// Slice Access
	// ------------
	fmt.Println("Slice Access:")

	sliceD := []int{10, 20, 30}

	// Accessing the values
	fmt.Println("sliceD[0] =", sliceD[0])
	fmt.Println("sliceD[1] =", sliceD[1])
	fmt.Println("sliceD[2] =", sliceD[2])

	// Changing the values
	sliceD[0] = 15
	sliceD[1] = 25
	sliceD[2] = 35

	fmt.Println("sliceD =", sliceD)
	fmt.Println()

	// Using Slices Helper Functions: slices
	// -------------------------------------
	fmt.Println("Using Slices Helper Functions:")

	sliceX := []int{1, 2, 3, 4, 5}
	sliceY := []int{1, 2, 3, 4, 5}
	sliceZ := []int{1, 2, 3, 4, 5, 6}
	sliceS := []string{"a", "b", "c"}

	fmt.Println("sliceX =", sliceX)
	fmt.Println("sliceY =", sliceY)
	fmt.Println("sliceZ =", sliceZ)
	fmt.Println("sliceS =", sliceS)

	fmt.Println("slices.Equal(sliceX, sliceY) ?", slices.Equal(sliceX, sliceY))
	fmt.Println("slices.Equal(sliceX, sliceZ) ?", slices.Equal(sliceX, sliceZ))
	// This one will not compile: Different types for each elements
	// fmt.Println("slices.Equal(sliceX, sliceS) ?", slices.Equal(sliceX, sliceS))

	fmt.Println()

	// Length of a Slice
	// -----------------
	fmt.Println("Length of a slice:")
	sliceF := []int{10, 20, 30, 40, 50}

	fmt.Println("sliceF =", sliceF, "length =", len(sliceF))

	fmt.Println()

	// Appending to a Slice
	// --------------------
	fmt.Println("Appending to a Slice:")
	var sliceG []int
	sliceG = append(sliceG, 100)
	// Appending to an existing slice
	sliceH := []int{1, 2, 3, 4, 5}
	// We can append more than 1 value at once
	sliceH = append(sliceH, 6, 7, 8, 9)
	// Appending one slice to another

	fmt.Println("sliceG =", sliceG)
	fmt.Println("sliceH =", sliceH)

	fmt.Println()

	// Extending a slice with another slice
	// ------------------------------------
	fmt.Println("Extending a Slice:")
	sliceH = append(sliceH, sliceG...)

	fmt.Println("sliceH =", sliceH)

	fmt.Println()

	// Appending and Slice Capacity
	// ----------------------------
	fmt.Println("Appending and Slice Capacity:")
	var xSlice []int
	fmt.Println(xSlice, len(xSlice), cap(xSlice))

	xSlice = append(xSlice, 10)
	fmt.Println(xSlice, len(xSlice), cap(xSlice))

	xSlice = append(xSlice, 20)
	fmt.Println(xSlice, len(xSlice), cap(xSlice))

	xSlice = append(xSlice, 30)
	fmt.Println(xSlice, len(xSlice), cap(xSlice))

	xSlice = append(xSlice, 40)
	fmt.Println(xSlice, len(xSlice), cap(xSlice))

	xSlice = append(xSlice, 50)
	fmt.Println(xSlice, len(xSlice), cap(xSlice))

	// Declaring a slice using make()
	// ------------------------------
	fmt.Println("Declaring a slice using make():")
	// Length: 5, Capacity: 5
	sliceI := make([]int, 5)
	// Length: 5, Capacity: 10
	sliceJ := make([]int, 5, 10)
	// Length: 0, Capacity: 10
	sliceK := make([]int, 0, 10)
	sliceK = append(sliceK, 1, 2, 3, 4, 5)

	fmt.Println("sliceI =", sliceI, "len(sliceI) =", len(sliceI), "cap(sliceI) =", cap(sliceI))
	fmt.Println("sliceJ =", sliceJ, "len(sliceJ) =", len(sliceJ), "cap(sliceJ) =", cap(sliceJ))
	fmt.Println("sliceK =", sliceK, "len(sliceK) =", len(sliceK), "cap(sliceK) =", cap(sliceK))

	fmt.Println()

	// Resetting slice using clear()
	// -----------------------------
	fmt.Println("Resetting slice using clear():")
	sliceL := []string{"first", "second", "third"}
	fmt.Println("Before:", sliceL, len(sliceL))
	clear(sliceL)
	fmt.Println("After:", sliceL, len(sliceL))

	fmt.Println()

	// Slicing slice
	// -------------
	fmt.Println("Slicing slice:")
	sliceM := []string{"a", "b", "c", "d"}
	subM1 := sliceM[:2]
	subM2 := sliceM[1:]
	subM3 := sliceM[1:3]
	subM4 := sliceM[:]

	fmt.Println("sliceM =", sliceM)
	fmt.Println("sliceM[:2] =", subM1)
	fmt.Println("sliceM[1:] =", subM2)
	fmt.Println("sliceM[1:3] =", subM3)
	fmt.Println("sliceM[:] =", subM4)

	fmt.Println()

	// Modifying one slice affects all others
	fmt.Println("Modifying one slice affects all others:")
	sliceM[1] = "y"
	subM1[0] = "x"
	subM2[1] = "z"

	fmt.Println("sliceM =", sliceM)
	fmt.Println("subM1 =", subM1)
	fmt.Println("subM2 =", subM2)
	fmt.Println("subM3 =", subM3)
	fmt.Println("subM4 =", subM4)

	fmt.Println()

	// Slice of slice with append()
	fmt.Println("Slice of slice with append()")
	fmt.Println("BEFORE:", "sliceM", "cap =", cap(sliceM), sliceM)
	fmt.Println("BEFORE:", "subM1", "cap =", cap(subM1), subM1)

	subM1 = append(subM1, "zz")

	fmt.Println("AFTER:", "sliceM", "cap =", cap(sliceM), sliceM)
	fmt.Println("AFTER:", "subM1", "cap =", cap(subM1), subM1)

	fmt.Println()

	// Full Slice Expressions
	fmt.Println("Full Slice Expressions:")
	sliceN := make([]string, 0, 5)
	sliceN = append(sliceN, "a", "b", "c", "d", "e")
	subN1 := sliceN[:2:2]
	subN2 := sliceN[2:4:4]

	fmt.Println("Capacities:", cap(sliceN), cap(subN1), cap(subN2))

	subN1 = append(subN1, "i", "j", "k")
	sliceN = append(sliceN, "x", "y")
	subN2 = append(subN2, "z", "zz")

	fmt.Println("sliceN =", sliceN)
	fmt.Println("subN1 =", subN1)
	fmt.Println("subN2 =", subN2)

	fmt.Println()

	// Slice Copy
	// ----------
	fmt.Println("Slicing copy:")
	sliceO := []int{1, 2, 3, 4}
	sliceP := make([]int, 4)

	fmt.Println("Before Copying:", sliceO, sliceP)
	num := copy(sliceP, sliceO)
	fmt.Println("After Copying:", sliceO, sliceP, "copied", num, "elements")

	fmt.Println()

	fmt.Println("Copying subset of slice:")
	sliceQ := make([]int, 2)
	num = copy(sliceQ, sliceO)
	fmt.Println("After Copying:", sliceO, sliceQ, "copied", num, "elements")

	fmt.Println()

	fmt.Println("Copying from the middle of slice:")
	sliceR := make([]int, 2)
	copy(sliceR, sliceO[2:])
	fmt.Println("After Copying:", sliceO, sliceR, "copied")

	fmt.Println()

	fmt.Println("Overlapping copies")
	num = copy(sliceO[:3], sliceO[1:])
	fmt.Println("After Copying:", sliceO, "copied", num, "elements")

	fmt.Println()

	fmt.Println("Using copy() with arrays:")
	sliceT := []int{1, 2, 3, 4} // Slice
	arrL := [4]int{5, 6, 7, 8}  //  Array
	sliceU := make([]int, 2)

	fmt.Println("Before copying:", sliceT, arrL, sliceU)
	copy(sliceU, arrL[:])
	fmt.Println("After copying from array:", sliceU, arrL)
	copy(arrL[:], sliceT)
	fmt.Println("After copying into array:", arrL, sliceT)

	fmt.Println()

	// Converting Array To Slice
	// -------------------------
	fmt.Println("Converting array to slice:")
	bArray := [4]int{100, 200, 300, 400}
	bSlice := bArray[:]
	fmt.Println("bArray =", bArray)
	fmt.Println("bSlice =", bSlice)

	fmt.Println()

	// Converting Slice to Array
	// -------------------------
	fmt.Println("Converting slice to array:")
	cSlice := []int{10, 20, 30, 40}
	cArray := [4]int(cSlice)
	smallArray := [2]int(cSlice)
	cSlice[0] = 100
	fmt.Println("cSlice =", cSlice)
	fmt.Println("cArray =", cArray)
	fmt.Println("smallArray =", smallArray)

	fmt.Println()

	// Example of Strings
	// ------------------
	fmt.Println("Example of Strings:")
	var strA string = "Hello there!"
	var strASub1 byte = strA[6]
	var strASub2 string = strA[4:7]
	var strASub3 string = strA[:5]
	var strASub4 string = strA[6:]
	fmt.Println("strA =", strA)
	fmt.Println("strASub1 =", strASub1)
	fmt.Println("strASub2 =", strASub2)
	fmt.Println("strASub3 =", strASub3)
	fmt.Println("strASub4 =", strASub4)

	fmt.Println()

	// String Code-point Issue
	// -----------------------
	fmt.Println("Example of Strings:")
	var strB string = "Hello 😊"
	var strBSub1 byte = strB[6]
	var strBSub2 string = strB[4:7]
	var strBSub3 string = strB[:5]
	var strBSub4 string = strB[6:]
	fmt.Println("strB =", strB)
	fmt.Println("strBSub1 =", strBSub1)
	fmt.Println("strBSub2 =", strBSub2)
	fmt.Println("strBSub3 =", strBSub3)
	fmt.Println("strBSub4 =", strBSub4)

	fmt.Println()

	// Length of String Issue
	// ----------------------
	fmt.Println("Length of String Issue:")
	var strC string = "Hello 😊"
	fmt.Println("strC =", strC)
	fmt.Println("len(strC) =", len(strC))

	fmt.Println()

	// Strings, runes, bytes type conversion
	// -------------------------------------
	fmt.Println("Strings, runes, bytes type conversion:")
	var a rune = 'a'
	stringA := string(a)
	var b byte = 'b'
	stringB := string(b)
	stringC := "Hello 😊"
	var bs []byte = []byte(stringC)
	var rs []rune = []rune(stringC)

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("stringA =", stringA)
	fmt.Println("stringB =", stringB)
	fmt.Println("stringC =", stringC)
	fmt.Println("bs =", bs)
	fmt.Println("rs =", rs)

	fmt.Println()

	// Declaring Maps
	// --------------
	fmt.Println("Declaring Maps:")
	var nilMap map[string]int
	notNilMap := map[string]int{}
	// Declaring a Map with values
	valuedMap := map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Mandarin"},
		"Lions":   {"Sarah", "Peter", "Mark"},
		"Kittens": {"Waldo", "Raul", "Alpha"},
	}
	// Creating a Map with make(): Specify the default size
	ages := make(map[int][]string, 10)

	fmt.Println("nilMap:", nilMap)
	fmt.Println("notNilMap:", notNilMap)
	fmt.Println("valuedMap:", valuedMap)
	fmt.Println("ages:", ages)
}

// FOR WINDOWS:
//  To Run:                 make run-win
//  To Build:               make build-win
//  To Run after Build:     .\bin\Composite.exe
//  Try Build + Run:        make try-win

// FOR LINUX:
//  To Run:                 make run
//  To Build:               make build
//  To Run after Build:     ./bin/Composite
//  Try Build + Run:        make try
