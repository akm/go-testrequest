# github.com/akm/go-testrequest

module: github.com/akm/go-testrequest
package: testrequest

## Overview

`go-testrequest` is a Go package that simplifies the creation and testing of HTTP requests. It provides a flexible and extensible way to build HTTP requests using a set of options and a factory pattern.

## Installation

```
go get github.com/akm/go-testrequest
```


## Usage

### Creating Requests

You can create HTTP requests using the provided functions such as `GET`, `POST`, `PUT`, etc. Each function accepts a list of options to configure the request.

```golang
import (
    "github.com/akm/go-testrequest"
)

req := testrequest.GET(
    testrequest.BaseUrl("http://example.com"),
    testrequest.Path("/users"),
    testrequest.Header("Authorization", "Bearer token"),
)
```

### Using Factories

Factories allow you to create requests with a set of default options. This is useful when you need to create multiple requests with the same base configuration.

```golang
factory := testrequest.NewFactory(testrequest.BaseUrl("http://example.com"))

req := factory.POST(
    testrequest.Path("/users"),
    testrequest.BodyString(`{"name":"John Doe"}`),
)
```

## Example Test

Here is an example of how to use go-testrequest in a test:

```golang
package testrequest

import (
    "net/http"
    "testing"

    "github.com/akm/go-testrequest"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestClientWithServer(t *testing.T) {
    testServer := startEchoServer(t)
    testServer.Start()
    defer testServer.Close()

    factory := testrequest.NewFactory(testrequest.BaseUrl(testServer.URL))

    client := &http.Client{}
    resp, err := client.Do(factory.GET(testrequest.Path("/foo")))
    require.NoError(t, err)
    defer resp.Body.Close()

    assert.Equal(t, http.StatusOK, resp.StatusCode)
}
```

This example demonstrates how to use `go-testrequest` to create and test HTTP requests in a Go test.  See [tests/client_test.go](./tests/client_test.go) for more details.

