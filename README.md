# svnparser

## Usage

```go
import csvparser

records, err := csvparser.Parse("test.csv")
if err != nil {
    println(err)
    return
}

for key, feilds := range records {
    println("key", key)
    for name, value := range feilds {
        println("name:", name, "value:", value)
    }
}
```

