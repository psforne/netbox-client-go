go-netbox 
=========

Package `netbox` provides an API 3.2 client for IPAM and DCIM service.

This package assumes you are using Netbox 3.2.

Using the client
================

The `github.com/psforne/netbox-client-go/netbox` package has some convenience functions for creating clients with the most common
```golang
package main

import (
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/psforne/netbox-client-go/netbox/client"
)

func main(){

    token := "xxx"
    netboxHost := "xxx"
    tlsoption := httptransport.TLSClientOptions{InsecureSkipVerify: true}
    tlsclient, err := httptransport.TLSClient(tlsoption)
    if err != nil {
        log.Fatal(err)
    }
    transport := httptransport.NewWithClient(netboxHost, client.DefaultBasePath, []string{"https"}, tlsclient)
    transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "Token "+token)
    c := client.New(transport, nil)
}

```

Go Module support
================

Go 1.16+

`go get github.com/psforne/netbox-client-go`


More complex client configuration
=================================

The client is generated using [go-swagger](https://github.com/go-swagger/go-swagger). This means the generated client
makes use of [github.com/go-openapi/runtime/client](https://godoc.org/github.com/go-openapi/runtime/client). If you need
a more complex configuration, it is probably possible with a combination of this generated client and the runtime
options.

The [godocs for the go-openapi/runtime/client module](https://godoc.org/github.com/go-openapi/runtime/client) explain
the client options in detail, including different authentication and debugging options. One thing I want to flag because
it is so useful: setting the `DEBUG` environment variable will dump all requests to standard out.

Regenerating the client
=======================

To regenerate the client with a new or different swagger schema, first clean the existing client, then replace
swagger.json and finally re-generate:
```
make clean
cp new_swagger_file.json swagger.json
make generate
```