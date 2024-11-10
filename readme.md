# Generics in GO
#### From Golang tutorial website


In this code, we will see the use of generics in golang to write more compact functions.

First we have writtent the `sumInt()` and `sumFloat()` functions without the use of generics.

```go
// Both functions have staticly defined parameters
func sumInts(m map[string]int64) int64 {}

func sumFloats(m map[string]float64) float64 {}
```

Now we would use generics to understand the difference and the use.

```go
func sumIntsOrFloats[K comparable, V Number](m map[K]V) V {}
```

In the above function, we decalre K as comparable so that we can use K as the key in the map variable. V on the other hand is declared as the `Number` interface type, so that it can understand whether the incoming value is of type `int64` or `float64`



Now, to run the code, open terminal and run

```
$ go run .
```


### Don't forget to give a star
## Happy Coding ^_^