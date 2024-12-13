<h1 align="center">conemu-progressbar for Go</h1>
<h6 align="center">Simple Go library for using ConEmu ProgressBar sequences (<code>OSC 9;4</code>)</h6>

# Installing
Run this:
```bash
$ go get github.com/ankddev/conemu-progressbar-go
```
Then, import in your code:
```go
import "github.com/ankddev/conemu-progressbar-go"
```
And enjoy!
# [Documentation](https://pkg.go.dev/github.com/ankddev/conemu-progressbar-go#section-documentation)
Functions:
* `ClearProgress` clears progress via sending sequence with 0 state
```go
ClearProgress()
```
* `SetIndeterminateProgress` sets indeterminate progress state
```go
SetIndeterminateProgress()
```
* `SetProgress` sets progress with default state
```go
err := SetProgress(80)
if err != nil {
    fmt.Println(err.Error())
}
```
Returns error if `progress` < `0` or > `100`
* `SetErrorProgress` sets progress with error state
```go
err := SetErrorProgress(50)
if err != nil {
    fmt.Println(err.Error())
}
```
Returns error if `progress` < `0` or > `100`
* `SetWarningProgress` sets progress with warning state
```go
err := SetWarningProgress(0)
if err != nil {
    fmt.Println(err.Error())
}
```
Returns error if `progress` < `0` or > `100`
# Contributing
* Fork repository
* Clone your fork
* Create changes
* Run tests with
```bash
$ go test
```
* Commit and push changes
* Create pull request
# See Also

- [codewars-api-rs](https://github.com/ankddev/codewars-api-rs) - Rust library for Codewars API
- [envfetch](https://github.com/ankddev/envfetch) - Lightweight crossplatform CLI tool for working with environment variables
- [terminal-go](https://github.com/ankddev/terminal-go) - Go library for working with ANSI/VT terminal sequences
- [zapret-discord-youtube](https://github.com/ankddev/zapret-discord-youtube) - Zapret build for Windows for fixing Discord and YouTube in Russia or othher services
