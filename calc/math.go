//Package calc This signifies that math.go belongs to calc package
//since this package is in a seperate directory under out module
//package declaration has nothing to do with module name
package calc

import "errors"

//Add - ...
func Add(numbers ...int) (error, int) {
	sum := 0

	if len(numbers) < 2 {
		return errors.New("provide more than 2"), sum
	} else {

		for _, num := range numbers {
			sum = sum + num
		}

		return nil, sum

	}

}
