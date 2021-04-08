# reflectx

[![codecov](https://codecov.io/gh/chyroc/reflectx/branch/master/graph/badge.svg?token=Z73T6YFF80)](https://codecov.io/gh/chyroc/reflectx)
[![go report card](https://goreportcard.com/badge/github.com/chyroc/reflectx "go report card")](https://goreportcard.com/report/github.com/chyroc/reflectx)
[![test status](https://github.com/chyroc/reflectx/actions/workflows/ci.yml/badge.svg)](https://github.com/chyroc/reflectx/actions)
[![Apache-2.0 license](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/chyroc/reflectx)

go reflect tool

## Install

```shell
go get github.com/chyroc/reflectx
```

## Usage

- check reflect.Value is some type

```go
// true
reflectx.IsUnsignedInt(reflect.ValueOf(uint(1)))

// false
reflectx.IsSignedInt(reflect.ValueOf(uint(1)))

// true
reflectx.IsInt(reflect.ValueOf(int(1)))
reflectx.IsInt(reflect.ValueOf(uint(1)))

// true
reflectx.IsBool(reflect.ValueOf(false))
```

- convert reflect.Value to some type
```go
// 1, nil
reflectx.ToInt64(reflect.ValueOf(uint(1)))

// 0, err
reflectx.ToInt64(reflect.ValueOf(false))

// true, nil
reflectx.ToBool(reflect.ValueOf(true))

// false, err
reflectx.ToBool(reflect.ValueOf(1))

// &true, nil
reflectx.ToPtr(reflect.ValueOf(true))

// &uint(1), err
reflectx.ToPtr(reflect.ValueOf(uint(1)))
```

## Dev

- codecov: https://github.com/codecov/example-go