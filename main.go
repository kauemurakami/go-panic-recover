package main

import "fmt"

func studentApproved(n1, n2 float64) bool {
	defer recoverExec()
	average := (n1 + n2) / 2
	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	//If it is equal to 6, your app goes into "panic"
	panic("A média é exatamente 6")
}

func recoverExec() {
	if r := recover(); r != nil {
		fmt.Println("Execution successfully recovered")
	}
}

func main() {
	fmt.Println(studentApproved(6, 7)) //outuput true
	fmt.Println(studentApproved(3, 7)) //outuput false
	// If the average is 6 and panic is called, the last print after it will not be executed
	fmt.Println(studentApproved(6, 6)) //outuput panic: The average is exactly 6 .....
	fmt.Println("Post execution")

}
