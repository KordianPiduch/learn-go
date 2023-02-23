# Functions

NOTE: In go functions there is no such thing as default value or keyword arguments.

Examples: 

``` Go 
package main

import "fmt"

func simpleSum(x ...int) (int, string) {
	// take multiply ints as positional arguments to sum
    // if multiple arguments is provided to function ...int should be at the end.
	sum := 0
	for _, item := range x {
		sum += item
	}
	return sum, "some text"
}

func simpleSum2(x []float64) float64 {
	// take slice of floats to sum
	sum := 0.0
	for _, item := range x {
		sum += item
	}
	return sum
}

func main() {
	y, s := simpleSum(1, 2, 3)
	fmt.Println(y, s)

	x := []int{5, 5, 5}
	z, _ := simpleSum(x...) // x... works as *args in Python
	fmt.Println(z)

	xx := []float64{4.0, 4.0}
	zz := simpleSum2(xx)
	fmt.Println(zz)
}
```

## Variadic functions

By using an ellipsis (...) before the type name of the last parameter, you can indicate that it takes zero or more of those parameters. In this case, we take zero or more ints.

``` Go
func add(args ...int) int { 
    total := 0
    for _, v := range args { 
        total += v
    }
    return total 
}
```
We can also pass a slice of ints by following the slice with an ellipsis:
``` Go
func main() {
    xs := []int{1,2,3} 
    fmt.Println(add(xs...))
}
```

## Clousers

A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
``` Go
func adder() func(int) {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main () {
    pos := adder()
    for i:=0; i<10; i++ {
        fmt.Println(pos(i))
    }
}
``` 

A function like this together with the nonlocal variables it references is known as a closure. In this case, increment and the variable x form the closure.

``` Go
func main() { 
    x:=0
    increment := func() int { 
        x++
        return x 
    }
    fmt.Println(increment())
    fmt.Println(increment())
}
```

Another way to use closure is by writing a function that returns another function, which when called, can generate a sequence of numbers. For example, here’s how we might generate all the even numbers:

``` Go 
func makeEvenGenerator() func() uint { 
    i := uint(0)
    return func() (ret uint) { 
        ret=i
        i+=2
        return
     } 
}
func main() {
    nextEven := makeEvenGenerator() 
    fmt.Println(nextEven()) // 0 
    fmt.Println(nextEven()) // 2 
    fmt.Println(nextEven()) // 4
}
```

