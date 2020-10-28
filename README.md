# Oto

Go driven RPC code generation tool for right now.

- 100% Go
- Describe services with Go interfaces (with structs, methods, comments, etc.)
- Generate server and client code
- Production ready templates (or copy and modify)

## This Fork

- CLI flag to output JSON for external utility consumption.

  I prefer dealing with JSON than having to reference oto's internal data structures.
  JSON is much easier to reference when writing templates.

  ```sh
  oto -json rpc.go > foo.json
  jq '.packageName' foo.json
  ```

- Changed comments metadata directive to `META(key):` like Go's `BUG(user):`.
  META value may span multiple lines stopping at next META or end of comment
  block. The value must be valid JSON.

  The existing convention of `key:` is how I usually mark sections in comments.
  Unfortunately, this is a BREAKING change and prevents merging to upstream project :(

- Added package level comments and metadata

  ```go
  // Package rpc comment.
  //
  // META(namespace): "github.com/mgutz/c"
  // META(imports): [
  //  "github.com/mgutz/a",
  //  "github.com/mgutz/b"
  // ]
  package rpc
  ```

## Templates

These templates are already being used in production.

- There are some [official Oto templates](https://github.com/pacedotdev/oto/tree/master/otohttp/templates)
- The [Pace CLI tool](https://github.com/pacedotdev/pace/blob/master/oto/cli.go.plush) is generated from an open-source CLI template

## Learn

![](oto-video-preview.jpg)

- VIDEO: [Mat Ryer gives an overview of Oto at the Belfast Gophers meetup](https://www.youtube.com/watch?feature=youtu.be&v=DUg4ZITwMys)

- BLOG: [How code generation wrote our API and CLI](https://pace.dev/blog/2020/07/27/how-code-generation-wrote-our-api-and-cli.html)

## Tutorial

Install the project:

```

go install github.com/pacedotdev/oto

```

Create a project folder, and write your service definition as a Go interface:

```go
// definitions/definitons.go
package definitions

// GreeterService makes nice greetings.
type GreeterService interface {
    // Greet makes a greeting.
    Greet(GreetRequest) GreetResponse
}

// GreetRequest is the request object for GreeterService.Greet.
type GreetRequest struct {
    // Name is the person to greet.
    // example: "Mat Ryer"
    Name string
}

// GreetResponse is the response object containing a
// person's greeting.
type GreetResponse struct {
    // Greeting is the greeting that was generated.
    // example: "Hello Mat Ryer"
    Greeting string
}
```

Download templates from otohttp

```bash
mkdir templates \
    && wget https://raw.githubusercontent.com/pacedotdev/oto/master/otohttp/templates/server.go.plush -q -O ./templates/server.go.plush \
    && wget https://raw.githubusercontent.com/pacedotdev/oto/master/otohttp/templates/client.js.plush -q -O ./templates/client.js.plush
```

Use the `oto` tool to generate a client and server:

```bash
mkdir generated
oto -template ./templates/server.go.plush \
    -out ./generated/oto.gen.go \
    -ignore Ignorer \
    -pkg generated \
    ./definitions
gofmt -w ./oto.gen.go ./oto.gen.go
oto -template ./templates/client.js.plush \
    -out ./generated/oto.gen.js \
    -ignore Ignorer \
    ./definitions
```

- Run `oto -help` for more information about these flags

Implement the service in Go:

```go
// greeter_service.go
package main

// GreeterService makes nice greetings.
type GreeterService struct{}

// Greet makes a greeting.
func (GreeterService) Greet(ctx context.Context, r GreetRequest) (*GreetResponse, error) {
    resp := &GreetResponse{
        Greeting: "Hello " + r.Name,
    }
    return resp, nil
}
```

Use the generated Go code to write a `main.go` that exposes the server:

```go
// main.go
package main

func main() {
    g := GreeterService{}
    server := otohttp.NewServer()
    generated.RegisterGreeterService(server, g)
    http.Handle("/oto/", server)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Use the generated client to access the service in JavaScript:

```javascript
import { GreeterService } from 'oto.gen.js';

const greeterService = new GreeterService();

greeterService
  .greet({
    name: 'Mat',
  })
  .then((response) => alert(response.greeting))
  .catch((e) => alert(e));
```

## Specifying additional template data

You can provide strings to your templates via the `-params` flag:

```bash
oto \
    -template ./templates/server.go.plush \
    -out ./oto.gen.go \
    -params "key1:value1,key2:value2" \
    ./path/to/definition
```

Within your templates, you may access these strings with `<%= params["key1"] %>`.

## Comment metadata

It's possible to include additional metadata for package, services, methods, objects, and fields
in the comments.

```go
// Thing does something.
//
// META(field): "value"
type Thing struct {
    //...
}
```

The `Metadata["field"]` value will be the string `value`.

- The value must be valid JSON (for strings, use quotes)

Examples are officially supported, but all data is available via the `Metadata` map fields.

## JSON Conversion

To convert definitions to JSON. Templates are not processed.

```sh
oto -J greeter_service.go |  jq '.services[0].name'

oto -J greeter_service.go > greeter_service.json
node generate.js greeter_service.json
```

### Examples

To provide an example value for a field, you may use the `example:` prefix line
in a comment.

```go
// GreetRequest is the request object for GreeterService.Greet.
type GreetRequest struct {
    // Name is the person to greet.
    // example: "Mat Ryer"
    Name string
}
```

- The example must be valid JSON

The example is extracted and made available via the `Field.Example` field.

## Contributions

Special thank you to:

- @mgutz - for struct tag support
- @sethcenterbar - for comment metadata support

![A PACE. project](pace-footer.png)
