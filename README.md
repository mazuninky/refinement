[![Build Status](https://travis-ci.com/mazuninky/blood-contracts-go.svg?branch=master)](https://travis-ci.com/mazuninky/blood-contracts-go)

# Refinement types

## Installation

1. Just go get it:

```sh
$ go get -u github.com/mazuninky/refinement
```

2. Import it in your code:

```go
import refinement "github.com/mazuninky/refinement"
```

## Refinement Data Type

### Create type

```go
numberType := refinement.MustNewRegexType(`[0-9]+`)
```

### Pack and unpack

```go
numberPack := numberType.Pack("45")
number, err := numberPack.Unpack()
if err != nil {
    panic(err)
}
println(number)
```

### Or

```go
emailType := refinement.MustNewRegexType(emailRegex)
phoneType := refinement.MustNewRegexType(phoneRegex)
loginType := phoneType.Or(emailType)
```

### Pipe

TODO

## License

MIT