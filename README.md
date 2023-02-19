# Golang Errors

Playing around the upcoming golang errors v2.0

https://go.googlesource.com/proposal/+/master/design/29934-error-values.md

The output would look like

```text
call A:
    main.main
        /Users/nam/work/go/go-errors/main.go:55
    call B
    main.callA
        /Users/nam/work/go/go-errors/main.go:62
    call external service
    main.callB
        /Users/nam/work/go/go-errors/main.go:69
    external service is down

```

## Prerequisites

* golang v1.20.1