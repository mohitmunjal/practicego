//Package calc This signifies that math.go belongs to calc package
//since this package is in a seperate directory under out module
//package declaration has nothing to do with module name
package calc

//Add - ...
func Add(numbers ...int) int {
	sum := 0

	for _, num := range numbers {
		sum = sum + num
	}

	return sum
}
