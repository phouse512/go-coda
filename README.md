## go-coda

This is a simple golang wrapper around the [Coda](https://coda.io) api, to
allow for automations and more advanced tooling to be built on top of it. For
example usage, see below.


#### usage

`go-coda` uses modules, and can be added to your project as shown below, then
by running `go test` or `go build` to have the module downloaded and added to
your `go.mod` file automatically.

```
import "phouse512/go-coda"

import "log"


func main() {
    codaClient := DefaultClient("sample-api-key")

    docs, err := codaClient.ListDocuments(codaClient.ListDocumentsPayload{})
    if err != nil {
        panic(err) 
    }

    log.Printf("All documents: %s", docs)
}

```



#### development

To work on the `go-coda` module locally, I recommend creating another test
module that imports your local `go-coda` module so that you can test your
changes. You can do that as follows:

```
$ go mod edit -replace github.com/phouse512/go-coda=/put/abs/path/to/local/module/here
```

Please note that that requires the test module to be initialized using `go mod
init`.

Feel free to open a PR against the `prod` branch once you are happy with your
changes.

