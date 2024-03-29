# tennvet

## Install

```sh
$ go install github.com/tenntenn/tennvet@latest
```

## How to use

```sh
$ go vet -vettool=`go env GOPATH`/bin/tennvet ./... 
```

If you set PATH to `go env GOPATH`/bin, you can use `which` command.

```sh
$ go vet -vettool=`which tennvet` ./... 
```

## Analyzers

* [zagane](https://github.com/gcpug/zagane/zagane)
* [ctxfield](https://github.com/gostaticanalysis/ctxfield)
* [dupimport](https://github.com/gostaticanalysis/dupimport)
* [elseless](https://github.com/gostaticanalysis/elseless)
* [forcetypeassert](https://github.com/gostaticanalysis/forcetypeassert)
* [importgroup](https://github.com/gostaticanalysis/importgroup)
* [innertypealias](https://github.com/gostaticanalysis/innertypealias)
* [loopdefer](https://github.com/gostaticanalysis/loopdefer)
* [nilerr](https://github.com/gostaticanalysis/nilerr)
* [noctor](https://github.com/gostaticanalysis/noctor)
* [nofmt](https://github.com/gostaticanalysis/nofmt)
* [notest](https://github.com/gostaticanalysis/notest)
* [sqlrows](https://github.com/gostaticanalysis/sqlrows)
* [testhelper](https://github.com/gostaticanalysis/testhelper)
* [typednil](https://github.com/gostaticanalysis/typednil)
* [typeswitch](https://github.com/gostaticanalysis/typeswitch)
* [unitconst](https://github.com/gostaticanalysis/unitconst)
* [unused](https://github.com/gostaticanalysis/unused)
* [wraperrfmt](https://github.com/gostaticanalysis/wraperrfmt)
* [zapvet](https://github.com/gostaticanalysis/zapvet)
* [debugcode](https://github.com/gostaticanalysis/debugcode)
