# Flow control

## 1. loops


### infinite loop
``` Go 
for {
    fmt.Println("This is inifinite loop")
}
```

### for-each loop: 
Iterate through items  

range iterates over elements for different data structures (not only arrays and slices)

for arrays and slices, range provides the index and the value for each element

``` Go 
var bookings = []string{}

for index, booking := range bookings {
    var fullName = strings.Fields(booking)  // split string on space
    var firstName = fullName[0]
    fmt.Println(firstName)
}
```

### standard for loop

``` Go 
for i:=1; i<=10; i++ {
    fmt.Println(i)
}
```
<br>
<br>

## 2. if statement

``` Go 
if <expression> {
    ...
} else if <expression> {
    ...
} else {
    ...
}
```

if statement works only with boolean conditions. 
empty string "" or value 0 in if statement result with error: `non-boolean condition in if statement`
<br><br>

## 3. switch statement
``` Go 
i := 2
fmt.Print("Write ", i, " as ")
switch i {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
case 3:
    fmt.Println("three")
default:
    fmt.Println("Printed in case of no matches")
}
```
``` Go 
city = "London"
switch city {
    case "Berlin":
        // some code here
    case "London", "Amsterdam"
        // some code here will execute for city == London or Amsterdam
    default:
        // some code if city is not one of above
}
```
