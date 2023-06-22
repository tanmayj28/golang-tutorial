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

## ENV agnostic:
- Go lets you build exe, linux executables, mac executables.
- Type: `go env`
  - Notice a var called `GOOS`, this can be `darwin`, `windows`, `linux` etc. and lets you build for diff systems.
  - Now, if you want to build your application in windows or exe format, just run this (on your mac maybe): `GOOS="windows" go build`

## Memory Management: (malloc & dealloc)
- In GoLang, memory management happens automatically, so you won't have to worry about it.
- `new()`:
  - allocate memory but not INIT.
  - get a memory address.
  - zeroed storage. (can't put data initially)
- `make()`:
  - allocate memory and INIT.
  - get a memory address.
  - non-zeroed storage. (can put data initially)

- GC happens automatically. Out of Scope or nil
  - It was rewritten after some initial issues.
  - Package to be used is called `runtime`.

- The `GOGC` variable sets the initial GC target percentage. A collection is triggered when the ratio of freshly allocated data to live data remaining after previous collection reaches this percentage. The default is `GOGC=100`. Setting `GOGC=off` disables the GC entirely. The runtime/debug pcakage's `SetGCPercent()` function allows changing this percentage at run time.
- There are a lot of functions available too.

## Pointers
- When you have multiple variables, functions and their invocations. These variables are passed in multiple functions, sometimes a copy of these variables gets passed and this can cause irrefularities if not handled correctly. Pointers help us, in fixing these issues correctly, by passing the actual address of the variable making sure we are dealing with correct values and variables.
So, a pointer is nothing but a direct reference to the memory address of a variable.
- Basic syntax would be:
```
myNumber := 23
var ptr *int
var ptrMyNumber = &myNumber
```

## Arrays:
Not too much used in GoLang instead we use slices. Refer to the 08myarray exercise for basics.

## Slices:
- Under the hood, these are actually Arrays. We use these more because they are far more powerful in GoLang for getting stuff done.
- We would be using a lot of slices in GoLang as they have many in-built functions.

## Maps:
- They are basically Hashes. When you're looping over maps, if you declare `key` as `_` it will be ignored by golang.
- For examples, checkout `10mymaps/` directory.

## Structs:
- In Golang we don't have classes, instead we use Structs.
  - There is no Inheritence in Golang.
  - No super or parent or child etc.

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

#### `time` package:
- The Time package in golang is a bit tricky, so understanding this via doc is important.
- In golang, if you want to format time, you need to format it by passing a date example with which you would like to format. And crazy part is, that date is actually a fixed value, `01-02-2006 15:04:05 Monday`.
- There are many functions in-built like:
  - sleep
  - ticker
  - etc.

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
