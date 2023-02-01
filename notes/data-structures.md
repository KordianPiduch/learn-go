# Basic data structures


## Arrays

``` Go 
var a [2]string
a[0] = "Hello"
a[1] = "World"
fmt.Println(a[0], a[1])
fmt.Println(a)

primes := [6]int{2, 3, 5, 7, 11, 13}
fmt.Println(primes)
```
<br><br>

## Slices

An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

``` Go
primes := [6]int{2, 3, 5, 7, 11, 13}

var s []int = primes[1:4]
fmt.Println(s)
```

``` Go
list := []string{}

list = append(list, "new item")
list = append(list, "another item")

fmt.Println(list)
```
<br><br>

## Maps

doc: https://go.dev/blog/maps  

This variable m is a map of string keys to int values:
``` Go 
var m map[string]int
```
Map types are reference types, like pointers or slices, and so the value of m above is nil; it doesn’t point to an initialized map. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don’t do that. To initialize a map, use the built in make function:
``` Go
m = make(map[string]int)
```

Examples: 
``` Go 
m := make(map[string]float64)
m["pi"] = 3.1416
```

Name of map is pointer. 
``` Go
a := map[string]int{
    "key1": 1,
    }

b := a

b["key2"] = 2

fmt.Println(a)  // map[key1:1 key2:2]
fmt.Println(b)  // map[key1:1 key2:2]
```
To initialize a map with some data, use a map literal:

``` Go 
exampleDict := map[string]int{
    "one": 1,
    "two": 2,
}

fmt.Println(exampleDict["one"])
```
The same syntax may be used to initialize an empty map, which is functionally identical to using the make function:
``` Go 
m = map[string]int{}
```

Map returns the zero value for the value type if key was not found. 
``` Go
dict := map[string]int{
		"a": 1,
	    }

fmt.Println(dict["b"])  // print out 0 (zero value for type int)
```
We can check for zero value by using this syntax:
``` Go 
dict := map[string]int{
		"a": 1,
	    }

item, ok := dict["a"]

fmt.Println(item, ok)  // item == 1, ok == true
```


<br><br>