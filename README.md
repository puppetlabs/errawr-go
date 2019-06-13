# errawr-go

This library contains bindings for
[errawr](https://github.com/puppetlabs/errawr-gen) to work in Go projects. It is
used heavily during the code generation process to produce statically-typed error libraries.

## Using this library

This library uses [Go modules](https://blog.golang.org/using-go-modules).

```console
$ go get github.com/puppetlabs/errawr-go@v2
```

## Encoding and decoding errors

The `encoding` package provides utility functions for creating envelopes to transfer errors across application boundaries.

### Display

To render an errawr error to a user using JSON, use the `ForDisplay` or `ForDisplayWithSensitivity` functions. The JSON serialization of the result of these functions conforms to the errawr specification.

A complete example of encoding and rendering an error to an HTTP response:

```go
import (
    "net/http"

    "github.com/puppetlabs/errawr-go/pkg/encoding"
    "github.com/puppetlabs/errawr-go/pkg/errawr"
)

func WriteError(w http.ResponseWriter, err errawr.Error) {
	status := http.StatusInternalServerError
	if hm, ok := err.Metadata().HTTP(); ok {
		status = hm.Status()

		for key, values := range hm.Headers() {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(encoding.ForDisplay(err)); err != nil {
		// Force this request to be abandoned.
		panic(err)
	}
}
```

If you are a client of an API that returns errawr errors as JSON, you can convert the errors back into an errawr object:

```go
func ReadError(data []byte) errawr.Error {
    var env encoding.ErrorDisplayEnvelope
    if err := json.Unmarshal(data, &env); err != nil {
        // ...
    }

    return env.AsError()
}
```

### Transit

If your application requires error serialization internally, such as across a shared memory interface or using gRPC, you can retain the complete structure of the error using the `ForTransit` function and the `AsError` method of the `ErrorTransitEnvelope`. Never use this API to present errors externally.

## Test utilities

When testing your code, it may be inconvenient to write real errawr definitions
for errors, especially in callbacks. You can use the `testutil` package for this:

```go
import (
    "testing"

    "github.com/puppetlabs/errawr-go/pkg/errawr"
    "github.com/puppetlabs/errawr-go/pkg/testutil"
)

func TestCallback(t *testing.T) {
    err := ExecuteWithCallback(func() errawr.Error {
        return testutil.NewStubError("fake_code")
    })
    if err == nil {
        t.Error("expected non-nil error")
    }
}
```

**Important:** Do not use the `testutil` package outside of tests. It defeats the purpose of errawr!
