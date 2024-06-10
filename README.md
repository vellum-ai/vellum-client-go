# Vellum Go Library

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-SDK%20generated%20by%20Fern-brightgreen)](https://buildwithfern.com/?utm_source=vellum-ai/vellum-client-go/readme)
![license badge](https://img.shields.io/github/license/vellum-ai/vellum-client-go)
[![go shield](https://img.shields.io/badge/go-docs-blue)](https://pkg.go.dev/github.com/vellum-ai/vellum-client-go)

The Vellum Go library provides convenient access to the Vellum API from Go.

# Requirements

This module requires Go version >= 1.18.

# Installation

Run the following command to use the Vellum Go library in your module:
```sh
go get github.com/vellum-ai/vellum-client-go
```

# Usage

```go
import vellumclient "github.com/vellum-ai/vellum-client-go/client"

client := vellumclient.NewClient(vellumclient.WithApiKey("<YOUR_AUTH_TOKEN>"))
```

## Generate Completion

```go
import (
  vellum       "github.com/vellum-ai/vellum-client-go"
  vellumclient "github.com/vellum-ai/vellum-client-go/client"
)

client := vellumclient.NewClient(vellumclient.WithApiKey("<YOUR_AUTH_TOKEN>"))
response, err := client.ExecutePrompt(
  context.TODO(),
  &vellum.ExecutePromptRequest{
    PromptDeploymentName: vellum.String("<your-deployment-name>"),
    Inputs: []*vellum.PromptDeploymentInputRequest{
      {
        Type: "String",
        String: &vellum.StringInputRequest{
          Name:  "<input_a>",
          Value: "Hello, world!",
        },
      },
    },
  },
)
```

## Timeouts

Setting a timeout for each individual request is as simple as using the standard
`context` library. Setting a one second timeout for an individual API call looks
like the following:

```go
ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
defer cancel()

response, err := client.ExecutePrompt(
  ctx,
  &vellum.ExecutePromptRequest{
    PromptDeploymentName: vellum.String("<your-deployment-name>"),
    Inputs: []*vellum.PromptDeploymentInputRequest{
      {
        Type: "String",
        String: &vellum.StringInputRequest{
          Name:  "<input_a>",
          Value: "Hello, world!",
        },
      },
    },
  },
)
```

## Client Options

A variety of client options are included to adapt the behavior of the library, which includes
configuring authorization tokens to be sent on every request, or providing your own instrumented
`*http.Client`. Both of these options are shown below:

```go
client := vellumclient.NewClient(
  vellumclient.WithApiKey("<YOUR_AUTH_TOKEN>"),
  vellumclient.WithHTTPClient(
    &http.Client{
      Timeout: 5 * time.Second,
    },
  ),
)
```

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

## Errors

Structured error types are returned from API calls that return non-success status codes. For example,
you can check if the error was due to a bad request (i.e. status code 400) with the following:

```go
response, err := client.ExecutePrompt(
  ctx,
  &vellum.ExecutePromptRequest{
    PromptDeploymentName: vellum.String("<invalid-name>"), // Updated to use an invalid deployment name
    Inputs: []*vellum.PromptDeploymentInputRequest{
      {
        Type: "String",
        String: &vellum.StringInputRequest{
          Name:  "<input_a>",
          Value: "Hello, world!",
        },
      },
    },
  },
)
if err != nil {
  if badRequestErr, ok := err.(*vellum.BadRequestError); ok {
    // Do something with the bad request error
    // Handle the bad request error specifically
  }
  return err // Ensure to return the error if present
}
```

These errors are also compatible with the `errors.Is` and `errors.As` APIs, so you can access the error
like so:

```go
response, err := client.Generate(
  ctx,
  &vellum.ExecutePromptRequest{
    PromptDeploymentName: vellum.String("<invalid-name>"), // Updated to use an invalid deployment name
    Inputs: []*vellum.PromptDeploymentInputRequest{


      {
        Type: "String",
        String: &vellum.StringInputRequest{
          Name:  "<input_a>",
          Value: "Hello, world!",
        },
      },
    },
  },
)
if err != nil {
  var badRequestErr *vellum.BadRequestError
  if errors.As(err, badRequestErr) {
    // Do something with the bad request ...
  }
  return err
}
```

If you'd like to wrap the errors with additional information and still retain the ability to access the type
with `errors.Is` and `errors.As`, you can use the `%w` directive:

```go
response, err := client.Generate(
  ctx,
  &vellum.ExecutePromptRequest{
    PromptDeploymentName: vellum.String("<invalid-name>"), // Updated to use an invalid deployment name
    Inputs: []*vellum.PromptDeploymentInputRequest{
      {
        Type: "String",
        String: &vellum.StringInputRequest{
          Name:  "<input_a>",
          Value: "Hello, world!",
        },
      },
    },
  },
)
if err != nil {
  return fmt.Errorf("failed to generate response: %w", err)
}
```

## Streaming

Calling any of Vellum's streaming APIs is easy. Simply create a new stream type and read
each message returned from the server until it's done:

```go
stream, err := client.ExecutePromptStream(
    context.TODO(),
    &vellum.ExecutePromptStreamRequest{
        PromptDeploymentName: vellum.String("<your-deployment-name>>"),
        Inputs: []*vellum.PromptDeploymentInputRequest{
            {
                Type: "String",
                String: &vellum.StringInputRequest{
                    Name:  "<input_a>",
                    Value: "Hello, world!",
                },
            },
        },
    },
)
if err != nil {
  return nil, err
}

// Make sure to close the stream when you're done reading.
// This is easily handled with defer.
defer stream.Close()

for {
  message, err := stream.Recv()
  if errors.Is(err, io.EOF) {
    // An io.EOF error means the server is done sending messages
    // and should be treated as a success.
    break
  }
  if err != nil {
    // The stream has encountered a non-recoverable error. Propagate the
    // error by simply returning the error like usual.
    return nil, err
  }
  // Do something with the message!
}
```

In summary, callers of the stream API use `stream.Recv()` to receive a new
message from the stream. The stream is complete when the `io.EOF` error is
returned, and if a non-`io.EOF` error is returned, it should be treated just
like any other non-`nil` error.

## Contributing

While we value open-source contributions to this SDK, most of this library is generated programmatically.

Please feel free to make contributions to any of the directories or files below:
```plaintext
tests/*
README.md
```

Any additions made to files beyond those directories and files above would have to be moved over to our generation code
(found in the separate [vellum-client-generator](https://github.com/vellum-ai/vellum-client-generator) repo),
otherwise they would be overwritten upon the next generated release. Feel free to open a PR as a proof of concept,
but know that we will not be able to merge it as-is. We suggest opening an issue first to discuss with us!