# prettybool [![Build Status](https://travis-ci.com/RaphaelPour/prettybool.svg?branch=master)](https://travis-ci.com/RaphaelPour/prettybool)
Pretty prints boolean values in various styles.

## Styles
Pretty Boolean styles (via `ListStyles()`):

| Style    | True     | False    |
| -------- | -------- | -------- |
| ok       | ok       | not ok   |
| yes      | yes      | no       |
| true     | true     | false    |
| 0        | 0        | 1        |
| passed   | passed   | failed   |
| positive | positive | negative |
| good     | good     | bad      |
| check    | ✅       | ❌       |
| enabled  | enabled  | disabled |
| on       | on       | off      |
| active   | active   | inactive |
| success  | success  | failure  |

## Example

```go
fmt.Printf("Tests: %s\n", styles.GetPrettyBool(true, "passed"))

// Tests: passed
```
