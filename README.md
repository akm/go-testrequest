# github.com/akm/reqb

![](https://img.shields.io/github/actions/workflow/status/akm/reqb/ci.yml)
![](https://img.shields.io/github/license/akm/reqb)


## Overview

`reqb` means `request builder`. The official pronunciation is 'rek-b'. It is a Go package that simplifies the creation and testing of HTTP requests. It provides a flexible and extensible way to build HTTP requests using a set of options and a factory pattern.

## Installation

```
go get github.com/akm/reqb
```


## Usage

### Creating Requests

You can create HTTP requests using the provided functions such as `GET`, `POST`, `PUT`, etc. Each function accepts a list of options to configure the request.

```golang
import (
    "github.com/akm/reqb"
)

req := reqb.GET(
    reqb.BaseUrl("http://example.com"),
    reqb.Path("/users"),
    reqb.Header("Authorization", "Bearer token"),
)
```

### Using Options type

`Options` allows you to create requests with a set of some options. This is useful when you need to create multiple requests with the same base configuration.

```golang
defaultOpts := reqb.Options{reqb.BaseUrl("http://example.com")}

req := defaultOpts.POST(
    reqb.Path("/users"),
    reqb.BodyString(`{"name":"John Doe"}`),
)
```

## Example Test

Here is an example of how to use reqb in a test:

```golang
package reqb

import (
    "net/http"
    "testing"

    "github.com/akm/reqb"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestClientWithServer(t *testing.T) {
    testServer := startEchoServer(t)
    testServer.Start()
    defer testServer.Close()

    factory := reqb.Options{reqb.BaseUrl(testServer.URL)}

    client := &http.Client{}
    resp, err := client.Do(factory.GET(reqb.Path("/foo")))
    require.NoError(t, err)
    defer resp.Body.Close()

    assert.Equal(t, http.StatusOK, resp.StatusCode)
}
```

This example demonstrates how to use `reqb` to create and test HTTP requests in a Go test.  See [tests/client_test.go](./tests/client_test.go) for more details.

