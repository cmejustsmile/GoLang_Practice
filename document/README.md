# Go Grammar
## Uncategorized
### Pointer
- Pointer Exists! (*, &)
- No Pointer Arithmetic

### struct
```
type Me struct{
    name string
    age int
}

func main(){
    who := Me{"Who", 100}
    fmt.Println()
}
```

### array(fixed size) : slice(dynamically-sized)
```
//array
var a [2]string
primes := [3]int{2,3,5}

// slice
var s []int = primes[2:3] // [lowerbound idx:higherbound idx], last one excluded

```
- slice references the array. It actually changes the value if modified.
- len(s): length, cap(s):# of element from the designated first element
- sliced slice's index is redifined([2:0]: 2nd idx-> 1st idx)
- empty slice = nil

```
// slice [0,0,0,0,0]
a := make([]int, 7)
a = append(a, 1)

board := [][]string{
	string{"_", "_", "_"},
	string{"_", "_", "_"},
	string{"_", "_", "_"},
}
	
```

### map
```
var m map[string]int
m["first"] = 1
fmt.Println(m["first"])

// map 
var m2 = map[string]int{
    "first" : 1,
    "second" : 2,
}

```

### iteration
```
// _ can replace i (skipping)
for i, v := range list {
    // i : idx
    // v : value in list
}

for
```

### Function
- functions are values!
```
wow := func(x int16) int16{
    return 2 * x
}

wow(4)

func compute(f func(int16) int16) int16{
    return f(5)
}

compute(wow)

```
- func can be return type!


### method
- method is a function with a special receiver argument
```
func (i struct) add2() int16{
    return i.x+2
}

i.add2()
```

### interface
- a set of method signatures
- requires method & struct 
[https://go-tour-kr.appspot.com/#53](https://go-tour-kr.appspot.com/#53)
- empty interface( interface{} ) as camelon

### type assertion
```
// does i has float64 value?
var i interface{} = "wow"
s, ok := i.(float64)
//panic. s has string
s :=i.(float64)
```


### package
- fmt
- io : Read
- image

### Goroutines
- new thread
```
go say("hi")
```

