## Golang:
- Install latest Golang from the website instructions.
- Setup proper Intellisense in your code editor for Go.

## Getting Started:
- Create a directory and run `go mod init hello`
  - This will create a `go.mod` file in your directory
  - `hello` will be your module name.
- Let's create your first program: `main.go`
```
pacakge main
import fmt
func main() {
  fmt.Println("Hello")
}
```
- Now we will run the file using: `go run main.go`
  - Note: This will only run the file and won't create any executables.

## Go Documentation
- Like every other langauge, Go also has a command line toolchain, and you can read more about commands using: `go help <command>`, `go help run` etc.
### `GOPATH`:
- It is an ENV variable.
- It has path for where to look for Go code.
- If the environment variable is unset, GOPATH defaults to a subdirectory named `go` in the user's home directory `$HOME/go` on Unix, `%USERPROFILE%\go` on Windows, unless that directory holds a Go distribution. Run `go env GOPATH` to see the current GOPATH.

### Go Packages:
- If you want to check out packages, head over to `pkg.go.com`.
- Let's explore some packages:

#### `buffio` package
- It is a buffer which can be used to read input from various devices and store them in a buffer.

#### `os` package

## Lexer and Types:
A lexer basically keeps check if the grammar / syntax of the language is properly followed or not.
- This Lexical Analysis happens before compilation.
- Therefore it is a good practcie to have Intellisense enabled for Go in your editor.
- For eg: In Go you don't need to write semicolons `;`. Lexer does it for us.

Types are helpful in determining the type of data stored in variable.
- Case sensitive, almost.
- V Type should be known in advance.
- Almost everything is a Type.
- Basic Types include: `String, Bool, Integer (uint8, uint64, int8, int64, uintptr), Float(float32, float64), Complex (complex numbers: iota)`
- Advanced Types: `Array, Slice, Maps, Struct. Pointers`
- And probably everything else like: `Functions, Channels, Mutex, . . .` etc.


## TTR:
- Packages have methods exposed and these usually start with capital letters.
  - Eg: `fmt.Println()`

## FAQs
- GoLang team rewrote their entire implmentation of their GC, it's now written completely in Go (earlier C).
## Related Links:
https://www.youtube.com/watch?v=62qGe9yhiJI&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=3&ab_channel=HiteshChoudhary
