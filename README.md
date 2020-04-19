# prettybool
Pretty prints boolean values in various styles.

## Styles
Pretty Boolean styles (via `styles.ListStyles()`):

| Style    | True     | False   |
| -------- | -------- | ------- |
| ok       | ok       | not ok  |
| yes      | yes      | no      |
| true     | true     | false   |
| 0        | 0        | 1       |
| passed   | passed   | failed  |
| positive | positive | negative|
| good     | good     | bad     |
| check    | ✅       | ❌      |


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
