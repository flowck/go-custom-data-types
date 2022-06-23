# Comparing types

## Comparable types

For simple data structures, go enables direct comparisons

```go
    type Person struct {
        FirtName string
		LastName string
    }
	
	person1 := Person{ FirstName: "John", LastName: "Doe" }
	person2 := Person{ FirstName: "John", LastName: "Doe" }
	
	if person1 == person2 {
	    fmt.Println("We match");	
    }   
```

## Equatable

## Custom method or function

## User keys in a map