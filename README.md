# energy [![Go Report Card](https://goreportcard.com/badge/github.com/liru/energy)](https://goreportcard.com/report/github.com/liru/energy) [![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/liru/energy/master/LICENSE) [![](https://godoc.org/github.com/liru/energy?status.svg)](http://godoc.org/github.com/liru/energy)
Energy provides a concurrent energy system, useful for games and other applications. 

Energy was heavily inspired by Heungsub Lee's [energy](https://github.com/sublee/energy) module for Python.

## Example

```go
e := energy.New(10,10,time.Second)
e.Use()
fmt.Println(e)
// <Energy 9/10>

e.UseEnergy(5)
fmt.Println(e)
// <Energy 4/10>

time.Sleep(time.Second)
fmt.Println(e)
// <Energy 5/10>

ok := e.UseEnergy(6)
if ok {
    fmt.Println("Do something")   
} else {
    fmt.Println("Not enough energy")
}
// "Not enough energy"
```

## Installation

`go get -u -v github.com/liru/energy`

## Todo

- Add a WaitForUse() function that blocks until there is energy available.
- Write tests.

## Suggestions

Use the `issues` tab provided by Github at the top of this project's page.

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D
