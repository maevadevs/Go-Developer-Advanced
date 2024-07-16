package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

// This is the main entry of the application.
func main() {
	// Headers
	fmt.Println()
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println()

	// Array Declarations
	// ------------------
	fmt.Println("Array Declaration:")
	fmt.Println("------------------")

	// Basic array declaration
	var arrA [3]int
	// Array declaration with Array literal
	arrB := [3]int{10, 20, 30}
	// Sparse array with specific non-zero elements
	arrC := [15]int{1, 5: 4, 6, 10, 100, 15}
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
	fmt.Println("-----------------")

	arrF := [...]int{1, 2, 3, 4, 5}
	arrG := [5]int{1, 2, 3, 4, 5}

	fmt.Println("arrF =", arrF)
	fmt.Println("arrG =", arrG)
	fmt.Println("arrF == arrG?", arrF == arrG)
	fmt.Println()

	// Multidimensional Array
	// ----------------------
	fmt.Println("Multidimensional Arrays:")
	fmt.Println("------------------------")

	// Array of array of integers: 5 x 5
	var arrH [5][5]int

	fmt.Println(arrH)
	fmt.Println()

	// Array Access
	// ------------
	fmt.Println("Array Access:")
	fmt.Println("-------------")

	arrI := [3]int{10, 20, 30}

	fmt.Println("arrI =", arrI)

	// Accessing the values
	fmt.Println("arrI[0] =", arrI[0])
	fmt.Println("arrI[1] =", arrI[1])
	fmt.Println("arrI[2] =", arrI[2])

	// Changing the array values
	arrI[0] = 15
	arrI[1] = 25
	arrI[2] = 35

	fmt.Println("After changing the array values:")
	fmt.Println("arrI =", arrI)
	fmt.Println()

	// Length of an array
	// ------------------
	fmt.Println("Length of An Array:")
	fmt.Println("-------------------")

	arrJ := [3]int{10, 20, 30}
	arrK := [...]int{15, 25, 35}

	fmt.Println("arrJ =", arrJ, "=> length:", len(arrJ))
	fmt.Println("arrK =", arrK, "=> length:", len(arrK))
	fmt.Println()

	// Slice Declarations
	// ------------------
	fmt.Println("Slice Declarations:")
	fmt.Println("-------------------")

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
	fmt.Println("------------------------")

	// Slice of slices of integers
	// Default value is a nil slice
	var sliceC [][]int

	// For better definition, we use make()
	// This initialize the slice with Zero-values instead of nil
	sliceC = make([][]int, 4)

	sliceC[0] = []int{1, 2, 3, 4}
	sliceC[1] = []int{5, 6, 7, 8}
	sliceC[2] = []int{9, 10, 11, 12}
	sliceC[3] = []int{13, 14, 15, 16}

	fmt.Println("sliceC =", sliceC)
	fmt.Println()

	// Slice Access
	// ------------
	fmt.Println("Slice Access:")
	fmt.Println("-------------")

	sliceD := []int{10, 20, 30}

	// Accessing the values
	fmt.Println("sliceD =", sliceD)
	fmt.Println("sliceD[0] =", sliceD[0])
	fmt.Println("sliceD[1] =", sliceD[1])
	fmt.Println("sliceD[2] =", sliceD[2])

	// Changing the values
	fmt.Println("Changing the values:")
	sliceD[0] = 15
	sliceD[1] = 25
	sliceD[2] = 35

	fmt.Println("sliceD =", sliceD)
	fmt.Println()

	// Using Slices Helper Functions: slices
	// -------------------------------------
	fmt.Println("Using Slices Helper Functions:")
	fmt.Println("------------------------------")

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
	fmt.Println("------------------")

	sliceF := []int{10, 20, 30, 40, 50}

	fmt.Println("sliceF =", sliceF, "=> length =", len(sliceF))
	fmt.Println()

	// Appending to a Slice
	// --------------------
	fmt.Println("Appending to a Slice:")
	fmt.Println("---------------------")

	// Appending to a nil slice
	var sliceG []int
	// Appending to an existing slice
	sliceH := []int{1, 2, 3, 4, 5}

	fmt.Println("Before: sliceG =", sliceG)
	fmt.Println("Before: sliceH =", sliceH)

	sliceG = append(sliceG, 100)
	// We can append more than 1 value at once
	sliceH = append(sliceH, 6, 7, 8, 9)
	// Appending one slice to another

	fmt.Println("After: sliceG =", sliceG)
	fmt.Println("After: sliceH =", sliceH)
	fmt.Println()

	// Extending a slice with another slice
	// ------------------------------------
	fmt.Println("Extending a Slice:")
	fmt.Println("------------------")

	sliceH = append(sliceH, sliceG...)

	fmt.Println("Extending sliceH with sliceG =", sliceH)
	fmt.Println()

	// Appending and Slice Capacity
	// ----------------------------
	fmt.Println("Appending and Slice Capacity:")
	fmt.Println("-----------------------------")

	var xSlice []int
	fmt.Println("Slice\t\t\tLen\tCap")
	fmt.Println(xSlice, "\t\t\t", len(xSlice), "\t", cap(xSlice))

	xSlice = append(xSlice, 10)
	fmt.Println(xSlice, "\t\t\t", len(xSlice), "\t", cap(xSlice))

	// We reach the current slice cap here
	xSlice = append(xSlice, 20)
	fmt.Println(xSlice, "\t\t", len(xSlice), "\t", cap(xSlice))

	// This one double the cap
	// Copy into a new slice
	xSlice = append(xSlice, 30)
	fmt.Println(xSlice, "\t\t", len(xSlice), "\t", cap(xSlice))

	// We reach the current slice cap here
	xSlice = append(xSlice, 40)
	fmt.Println(xSlice, "\t\t", len(xSlice), "\t", cap(xSlice))

	// This one double the cap
	// Copy into a new slice
	xSlice = append(xSlice, 50)
	fmt.Println(xSlice, "\t", len(xSlice), "\t", cap(xSlice))
	fmt.Println()

	// Declaring a slice using make()
	// ------------------------------
	fmt.Println("Declaring a slice using make():")
	fmt.Println("-------------------------------")

	// Length: 5, Capacity: 5
	sliceI := make([]int, 5)
	// Length: 5, Capacity: 10
	sliceJ := make([]int, 5, 10)
	// Length: 0, Capacity: 10
	sliceK := make([]int, 0, 10)

	fmt.Println("Before Append:")
	fmt.Println("\tsliceI =", sliceI, "len(sliceI) =", len(sliceI), "cap(sliceI) =", cap(sliceI))
	fmt.Println("\tsliceJ =", sliceJ, "len(sliceJ) =", len(sliceJ), "cap(sliceJ) =", cap(sliceJ))
	fmt.Println("\tsliceK =", sliceK, "len(sliceK) =", len(sliceK), "cap(sliceK) =", cap(sliceK))

	sliceK = append(sliceK, 1, 2, 3, 4, 5)

	fmt.Println("After Append to sliceK:")
	fmt.Println("\tsliceK =", sliceK, "len(sliceK) =", len(sliceK), "cap(sliceK) =", cap(sliceK))
	fmt.Println()

	// Resetting slice using clear()
	// -----------------------------
	fmt.Println("Resetting slice using clear():")
	fmt.Println("------------------------------")

	sliceL := []string{"first", "second", "third"}
	sliceLInt := []int{100, 200, 300}
	fmt.Println("Before:")
	fmt.Println("\tsliceL =", sliceL, "len =", len(sliceL), "cap =", cap(sliceL))
	fmt.Println("\tsliceLInt =", sliceLInt, "len =", len(sliceLInt), "cap =", cap(sliceLInt))
	clear(sliceL)
	clear(sliceLInt)
	fmt.Println("After:")
	fmt.Println("\tsliceL =", sliceL, "len =", len(sliceL), "cap =", cap(sliceL))
	fmt.Println("\tsliceLInt =", sliceLInt, "len =", len(sliceLInt), "cap =", cap(sliceLInt))
	fmt.Println()

	// Subslicing slice
	// ----------------
	fmt.Println("Subslicing slice:")
	fmt.Println("-----------------")

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
	// --------------------------------------
	fmt.Println("Modifying one subslice affects all others:")
	fmt.Println("------------------------------------------")

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
	// ----------------------------
	fmt.Println("Slice of slice with append():")
	fmt.Println("-----------------------------")

	fmt.Println("BEFORE:")
	fmt.Println("\tsliceM", "cap =", cap(sliceM), sliceM)
	fmt.Println("\tsubM1", "cap =", cap(subM1), subM1)

	subM1 = append(subM1, "zz")

	fmt.Println("AFTER:")
	fmt.Println("\tsliceM", "cap =", cap(sliceM), sliceM)
	fmt.Println("\tsubM1", "cap =", cap(subM1), subM1)
	fmt.Println()

	// Full Slice Expressions
	// ----------------------
	fmt.Println("Full Slice Expressions:")
	fmt.Println("-----------------------")

	sliceN := make([]string, 0, 5)
	sliceN = append(sliceN, "a", "b", "c", "d", "e")
	subN1 := sliceN[1:4:4]
	subN2 := sliceN[2:4:4]

	fmt.Println("sliceN =", sliceN)
	fmt.Println("subN1 =", subN1)
	fmt.Println("subN2 =", subN2)
	fmt.Println("Capacities:", cap(sliceN), cap(subN1), cap(subN2))

	subN1 = append(subN1, "i", "j", "k")
	sliceN = append(sliceN, "x", "y")
	subN2 = append(subN2, "z", "zz")

	fmt.Println("Operations:")
	fmt.Println("\tappend(subN1, \"i\", \"j\", \"k\")")
	fmt.Println("\tappend(sliceN, \"x\", \"y\")")
	fmt.Println("\tappend(subN2, \"z\", \"zz\")")

	fmt.Println("sliceN =", sliceN)
	fmt.Println("subN1 =", subN1)
	fmt.Println("subN2 =", subN2)
	fmt.Println()

	// Slicing Copy
	// ------------
	fmt.Println("Slicing copy:")
	fmt.Println("-------------")

	sliceO := []int{1, 2, 3, 4, 5}
	sliceP := make([]int, 10)

	fmt.Println("Before Copying:", "src =", sliceO, "dest =", sliceP)
	num := copy(sliceP, sliceO)
	fmt.Println("After Copying:", "src =", sliceO, "dest =", sliceP, "copied", num, "elements")
	fmt.Println()

	// Copying subset of slice
	// -----------------------
	fmt.Println("Copying subset of slice:")
	fmt.Println("------------------------")

	sliceQ := make([]int, 2)
	num = copy(sliceQ, sliceO)
	fmt.Println("After Copying:", "src =", sliceO, "dest =", sliceQ, "copied", num, "elements")
	fmt.Println()

	// Copying subset from the middle of slice
	// ---------------------------------------
	fmt.Println("Copying subset from the middle of slice:")
	fmt.Println("----------------------------------------")

	// Length of the destination slice can be derived from the subslice
	// If the return value is not needed, ignore
	sliceR := make([]int, len(sliceO[2:]))
	copy(sliceR, sliceO[2:])

	fmt.Println("After Copying:", "src =", sliceO, "dest =", sliceR)
	fmt.Println()

	// Overlapping copies
	// ------------------
	fmt.Println("Overlapping copies:")
	fmt.Println("-------------------")

	num = copy(sliceO[:3], sliceO[1:])
	fmt.Println("After Copying:", sliceO, "copied", num, "elements")
	fmt.Println()

	// Using copy() with arrays
	// ------------------------
	fmt.Println("Using copy() with arrays:")
	fmt.Println("-------------------------")

	sliceT := []int{1, 2, 3, 4} // Slice
	arrL := [4]int{5, 6, 7, 8}  // Array
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
	fmt.Println("--------------------------")

	bArray := [4]int{100, 200, 300, 400}
	fmt.Println("bArray Before =", bArray)

	bSlice := make([]int, len(bArray[:]))
	copy(bSlice, bArray[:])
	bSlice[0] = 10000 // No memory-sharing after using copy()
	bArray[3] = 40000 // No memory-sharing after using copy()

	fmt.Println("bArray =", bArray)
	fmt.Println("bSlice =", bSlice)
	fmt.Println()

	// Converting Slice to Array
	// -------------------------
	fmt.Println("Converting slice to array:")
	fmt.Println("--------------------------")

	cSlice := []int{10, 20, 30, 40}
	fmt.Println("cSlice Before =", cSlice)

	cArray := [4]int(cSlice)
	smallArray := [2]int(cSlice)
	cSlice[0] = 10000 // No memory-sharing
	cArray[3] = 40000 // No memory-sharing
	fmt.Println("cSlice =", cSlice)
	fmt.Println("cArray =", cArray)
	fmt.Println("smallArray =", smallArray)
	fmt.Println()

	// Example of Strings
	// ------------------
	fmt.Println("Example of Strings:")
	fmt.Println("-------------------")

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
	fmt.Println("Example of Strings With Code-Point Issues:")
	fmt.Println("------------------------------------------")

	var strB string = "Hello 😊!" // 😊 takes 4 bytes
	var strBSub1 byte = strB[6]
	var strBSub2 string = strB[4:7]
	var strBSub3 string = strB[:5]
	var strBSub4 string = strB[6:]

	fmt.Println("strB =", strB, "len =", len(strB)) // len() returns the number of bytes
	fmt.Println("strBSub1 =", strBSub1)
	fmt.Println("strBSub2 =", strBSub2)
	fmt.Println("strBSub3 =", strBSub3)
	fmt.Println("strBSub4 =", strBSub4)
	fmt.Println()

	// Length of String Issue
	// ----------------------
	fmt.Println("Length of String Issue:")
	fmt.Println("-----------------------")

	var strC string = "Hello 😊!"
	fmt.Println("strC =", strC)
	// len() returns the number of bytes, not runes
	fmt.Println("len(strC) =", len(strC))
	fmt.Println()

	// Strings, runes, bytes type conversion
	// -------------------------------------
	fmt.Println("Strings, runes, bytes type conversion:")
	fmt.Println("--------------------------------------")

	var charA rune = 'a'
	var charB byte = 'b'
	stringA := string(charA)
	stringB := string(charB)
	stringC := "Hello 😊!"
	var bytesStringC []byte = []byte(stringC)
	var runesStringC []rune = []rune(stringC)

	fmt.Println("charA =", charA)
	fmt.Println("charB =", charB)
	fmt.Println("stringA =", stringA)
	fmt.Println("stringB =", stringB)
	fmt.Println("stringC =", stringC)
	fmt.Println("bytesStringC =", bytesStringC)
	fmt.Println("runesStringC =", runesStringC)
	fmt.Println()

	// Declaring Maps
	// --------------
	fmt.Println("Declaring Maps:")
	fmt.Println("---------------")

	// nil Map
	var nilMap map[string]int
	// Map Literal (not nil)
	notNilMap := map[string]int{}
	// Declaring a Map with map literal and key-value pairs
	valuedMap := map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Mandarin"},
		"Lions":   {"Sarah", "Peter", "Mark"},
		"Kittens": {"Waldo", "Raul", "Alpha"},
	}
	// Creating a Map with make(): Similar to using Map Literal
	// But specify the default size
	agesMap := make(map[int][]string, 10)

	fmt.Println("nilMap:", nilMap)
	fmt.Println("notNilMap:", notNilMap)
	fmt.Println("valuedMap:", valuedMap)
	fmt.Println("agesMap:", agesMap)
	fmt.Println()

	// Reading & Writing to Maps
	// -------------------------
	fmt.Println("Reading and Writing to Maps:")
	fmt.Println("----------------------------")

	totalWins := map[string]int{}

	fmt.Println("Before:")
	fmt.Println("totalWins =", totalWins)

	// Writing to a Map
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	totalWins["Kittens"]++
	totalWins["Lions"] = 3

	// Reading a Map
	fmt.Println("After:")
	fmt.Println("totalWins[\"Orcas\"] =", totalWins["Orcas"])
	fmt.Println("totalWins[\"Lions\"] =", totalWins["Lions"])
	fmt.Println("totalWins[\"Kittens\"] =", totalWins["Kittens"])
	fmt.Println()

	// Comma-Ok Idiom
	// --------------
	fmt.Println("Comma-Ok Idiom:")
	fmt.Println("---------------")

	greetMap := map[string]int{
		"hello": 5,
		"world": 0,
	}
	fmt.Println(greetMap)

	// In the map with value 5
	val, ok := greetMap["hello"]
	fmt.Println("hello =>", val, ok)

	// In the map with value 0
	val, ok = greetMap["world"]
	fmt.Println("world =>", val, ok)

	// Not in the map, default value 0
	val, ok = greetMap["hi"]
	fmt.Println("hi =>", val, ok)
	fmt.Println()

	// Deleting From Maps
	// ------------------
	fmt.Println("Deleting From Maps:")
	fmt.Println("-------------------")

	delMap := map[string]int{
		"hello": 5,
		"world": 10,
	}

	fmt.Println("Before Delete:", delMap)
	delete(delMap, "hello")
	delete(delMap, "unknownKey") // Nothing happens
	fmt.Println("After Delete:", delMap)
	fmt.Println()

	// Clearing a Map
	// --------------
	fmt.Println("Clearing a Map:")
	fmt.Println("---------------")

	clearMap := map[string]int{
		"hello": 5,
		"world": 10,
	}

	fmt.Println("Before Clear:", clearMap, "len =", len(clearMap))
	clear(clearMap)
	fmt.Println("After Clear:", clearMap, "len =", len(clearMap))
	fmt.Println()

	// Comparing 2 maps
	// ----------------
	fmt.Println("Comparing 2 maps:")
	fmt.Println("-----------------")

	mp1 := map[string]int{
		"hello": 5,
		"world": 5,
	}
	mp2 := map[string]int{
		"hello": 5,
		"world": 5,
	}
	mp3 := map[string]int{
		"hello": 6,
		"world": 5,
	}

	fmt.Println("mp1 =", mp1)
	fmt.Println("mp2 =", mp2)
	fmt.Println("mp3 =", mp3)
	fmt.Println("maps.Equal(mp1, mp2): ", maps.Equal(mp1, mp2))
	fmt.Println("maps.Equal(mp1, mp3): ", maps.Equal(mp1, mp3))
	fmt.Println()

	// Using map[KeyType]bool as Set
	// -----------------------------
	fmt.Println("Using map[KeyType]bool as Set:")
	fmt.Println("------------------------------")

	intSet := map[int]bool{}
	intSet2 := map[int]string{}
	vals := []int{5, 10, 2, 1, 5, 8, 7, 3, 9, 8, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
		intSet2[v] = ""
	}

	fmt.Println("vals =", vals, "len =", len(vals))
	fmt.Println("intSet =", intSet, "len =", len(intSet))
	fmt.Println("intSet2 =", intSet2, "len =", len(intSet2))
	fmt.Println()

	// Checking for element in the set
	// -------------------------------
	fmt.Println("Checking for element in the set:")
	fmt.Println("--------------------------------")

	// Using Boolean as Value
	inIntSet := false
	if intSet[5] {
		inIntSet = true
	}
	fmt.Println("5 in intSet?", inIntSet)

	// Using the Comma-Ok Idiom
	_, inIntSet2 := intSet2[2000]
	fmt.Println("2000 in intSet2?", inIntSet2)
	fmt.Println()

	// Struct Definition and Declaration
	// ---------------------------------
	fmt.Println("Struct Definition and Declaration:")
	fmt.Println("----------------------------------")

	// Example of defining a struct
	type Person struct {
		fName   string
		lName   string
		dob     string
		favNum  int
		isAdult bool
	}

	// Declaring a struct with zero-values
	var fred Person
	fmt.Println("fred =", fred)

	// Declaring a struct with struct literal (zero-values)
	bob := Person{}
	fmt.Println("bob =", bob)

	// Declaring a struct with struct literal and values
	julia := Person{
		"julia",
		"smith",
		"1969-01-01",
		77,
		true,
	}
	fmt.Println("julia =", julia)

	// Declaring a struct with struct literal and values, with named-fields
	john := Person{
		dob:   "2023-01-01",
		lName: "smith",
		fName: "john",
	}
	fmt.Println("john =", john)
	fmt.Println()

	// Struct field access
	// -------------------
	fmt.Println("Struct field access:")
	fmt.Println("--------------------")

	bob.fName = "Bob"
	fmt.Println("bob.fName =", bob.fName)
	fmt.Println()

	// Anonymous Struct
	// ----------------
	fmt.Println("Anonymous Struct:")
	fmt.Println("-----------------")

	pet := struct {
		// Struct Definition
		name  string
		kind  string
		breed string
	}{
		// Struct value
		name:  "fido",
		kind:  "dog",
		breed: "golden retriever",
	}

	fmt.Println("pet.name:", pet.name)
	fmt.Println("pet.kind:", pet.kind)
	fmt.Println("pet.breed:", pet.breed)
	fmt.Println()

	// Comparing Against Anonymous Struct
	// ----------------------------------
	fmt.Println("Comparing Against Anonymous Struct:")
	fmt.Println("-----------------------------------")

	type firstPerson struct {
		name string
		age  int
	}

	f := firstPerson{
		name: "Bob",
		age:  50,
	}

	g := struct {
		name string
		age  int
	}{
		name: "Bob",
		age:  50,
	}

	// Comparing
	fmt.Println("f == g?", f == g)

	// Using with each other as equivalent
	f = g
	fmt.Println("f =", f)
	fmt.Println()

	fmt.Println(strings.Repeat("-", 100))
	fmt.Println()
}

// AVAILABLE COMMANDS
// ------------------
//  make                Default to `make try`
//  make fmt            Format all source files
//  make vet            Verify any possible errors
//  make build          Build module
//  make build-release  Build module for release, strip symbols
//  make run            Build module then run
//  make run-release    Build module for release then run
//  make try            Build module, run, then remove built binary
