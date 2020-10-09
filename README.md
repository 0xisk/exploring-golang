# Exploring Golang
A repository focused on learning and practicing Golang

<p align='center'>
    <img src='golang-logo.png' alt='screenshot' />
</p>

## Go (golang)

1. Strong and statically typed
2. Excellent community
3. Key features
   - Simplicity
   - Fast compile times
   - Garbage collected
   - Built-in concurrency
   - Compile to standalone binaries

## Variables
### Variable declaration:
- var foo int
- var foo int = 42
- foo := 42

### Cant redeclare variables, but can shadow them.

### All variables must be used.

### Visibility:
  - lower case first letter for package scope
  - upper case first letter to export
  - no private scope

### Naming conventions:
- Pascal or camelCase
    - Capitalize acronyms (HTTP, URL)
- As short as reasonable
    - Longer names for longer lives

### Type conversions:
- DestinationType(variable)
- Use (strconv) package for string

## Resources

- [golang.org](https://golang.org/)
- [golang docs](https://golang.org/doc/)