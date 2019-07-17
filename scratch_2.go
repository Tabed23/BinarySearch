package main

import (
	"fmt"
)


func main() {

	 array := [...]int { 1,2,3,4,5,6,7,8,9,10}
	if BinarySearch(array, 15) {
		fmt.Println("Is in the Array")
	} else {
		fmt.Println("Is not in tha Array")
	}

}

func BinarySearch(A[10] int, target int)bool  {

	var start int = A[0]
	var end int = A[len(A)-1]
	var mid int
	var found bool =  false
	for start < end{

		mid =  (start + end)/ 2
		if A[mid] > target{
			end = mid -1
		}
		if A[mid] < target{
			start = mid +1
		}
		if A[mid] ==  target{
			found = true
			break
		}
	}

	if found {
		return true
	}
	return false

}
