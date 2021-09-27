# assert-require

tl;dr — `assert` keeps running through the rest of the test, while `require`
will exit on the first failure.

In this module:

```
❯ go test -v ./
=== RUN   TestAssert
    main_test.go:11: 
                Error Trace:    main_test.go:11
                Error:          Not equal: 
                                expected: 1
                                actual  : 2
                Test:           TestAssert
    main_test.go:12: 
                Error Trace:    main_test.go:12
                Error:          Not equal: 
                                expected: 2
                                actual  : 3
                Test:           TestAssert
--- FAIL: TestAssert (0.00s)
=== RUN   TestRequire
    main_test.go:17: 
                Error Trace:    main_test.go:17
                Error:          Not equal: 
                                expected: 1
                                actual  : 2
                Test:           TestRequire
--- FAIL: TestRequire (0.00s)
FAIL
FAIL    ajm188.scratchpad/assert-require        0.531s
FAIL
```
