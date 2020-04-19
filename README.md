# prettybool
Pretty prints boolean values in various styles.

## Example

```go
pretty, err := styles.GetPrettyBool(true, "passed")

if err != nil {
    fmt.Println("Style does not exist")
    return
}

fmt.Printf("Tests: %s\n", pretty)

// Tests: passsed
```


