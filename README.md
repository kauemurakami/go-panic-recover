[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-named-return-functions/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-named-return-functions/blob/main/README.md)  
go version 1.22.1  

## Panic e Recover
Widely used with ```defer```, let's use the example that calculates the average of grades, but the average can never be 6, if it is 6 our program will go into "panic" using the clause ```panic```, what ```panic``` does is interrupt the normal flow of your program, stopping everything, stopping running everything and "panicing", a term used in go.  
Basically it will call all functions that have ```defer```, if you don't recover the function with ```recover```, which we will see later, your program dies.  

### Starting the project
Create a ```go-panic-recover``` directory with a ```main.go``` file containing:
```go
package main

func studentApproved(n1, n2 float64) bool {
	average := (n1 + n2) / 2
	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}
  //If it is equal to 6, your app goes into "panic"
	panic("A média é exatamente 6")
}

func main() {
fmt.Println(studentApproved(6, 7)) //outuput true
	fmt.Println(studentApproved(3, 7)) //outuput false
	// If the average is 6 and panic is called, the last print after it will not be executed
	fmt.Println(studentApproved(6, 6)) //outuput panic: The average is exactly 6 ..... errors
	fmt.Println("Pós execução")
}
```
Here we see that the program panicked, because we didn't expect the number 6 as the average, and the program doesn't know what to do, unlike ```error```, which you can continue and handle, ```panic``` kills the execution of the program, so if you have the ```panic``` clause and don't have ```recover```, your program dies.  
But there is a way to recover your program if it goes into ```panic```, which is with the ```recover``` clause, example:  
```go
.....
func studentApproved(n1, n2 float64) bool {
	defer recoverExec() //<<
	average := (n1 + n2) / 2
	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	//If it is equal to 6, your app goes into "panic"
	panic("A média é exatamente 6")
}
// function to recover project execution
func recoverExec() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func main() {
	fmt.Println(studentApproved(6, 7)) //outuput true
	fmt.Println(studentApproved(3, 7)) //outuput false
	// If the average is 6 and panic is called, the last print after it will not be executed
	fmt.Println(studentApproved(6, 6)) //outuput panic is ignored and we resume with recover
	fmt.Println("Pós execução")
}
```
We then saw that the ```recover()``` function can recover the execution of our program.
Understanding better by adding the ```defer``` clause when calling the ```studentApproved(...)``` function. When ```panic``` occurs, we always call the defer functions, which in our case is ```recoverExec()```, in it we check whether ```recover()``` managed to recover something, so it is != nil in this case, if it returns nil we cannot recover.  

