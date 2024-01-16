# gotls

## Overview

`gotls` is a Go package that extends crypto/tls by offering a comprehensive range of cipher suites. It is designed for developers who require advanced TLS configurations and a broader selection of cipher suites in their Go applications.

All tls information is fetched from `https://www.iana.org/assignments/tls-parameters` (Internet Assigned Numbers Authority)

## Installation

To install `gotls`, use the following go get command: `go get github.com/TechMDW/gotls`

## Usage

Below is a basic example of how to use `gotls` in your Go application:

```go
package main

import (
    "github.com/TechMDW/gotls"
    "fmt"
)

func main() {
    // Example: Fetching a cipher suite by name
    {
        cipherSuite := gotls.GetByName("TLS_RSA_WITH_AES_128_GCM_SHA256")
        if cipherSuite != nil {
            fmt.Println("Cipher Suite Name:", cipherSuite.Name)
            fmt.Println("Cipher Suite Code:", cipherSuite.Code)
        }
    }

    {
        cipherSuite := gotls.Get(156)
        if cipherSuite != nil {
            fmt.Println("Cipher Suite Name:", cipherSuite.Name)
            fmt.Println("Cipher Suite Code:", cipherSuite.Code)
        }
    }
}
```

## License

This project is licensed under the Apache License 2.0.
