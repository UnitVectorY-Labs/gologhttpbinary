[![GitHub release](https://img.shields.io/github/release/UnitVectorY-Labs/gologhttpbinary.svg)](https://github.com/UnitVectorY-Labs/gologhttpbinary/releases/latest) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![Active](https://img.shields.io/badge/Status-Active-green)](https://guide.unitvectorylabs.com/bestpractices/status/#active) [![Go Report Card](https://goreportcard.com/badge/github.com/UnitVectorY-Labs/gologhttpbinary)](https://goreportcard.com/report/github.com/UnitVectorY-Labs/gologhttpbinary)

# gologhttpbinary

A lightweight HTTP server that logs HTTP requests, including the body encoded in base64, enabling logging of binary payloads.

## Purpose

This application captures all incoming HTTP requests and logs the path, headers, and the body of the request as a base64 encoded string. This is useful for debugging and logging binary payloads in a container environment.

**Why use this?** It’s a straightforward tool for debugging, testing, and development, providing an easy way log requests to an HTTP server for inspection. The reason this was created was to log progobufs that were being sent to a server for debugging and testing as other tools were not capturing the binary data therefore logging the request as base64 as a solution.

**Should I run this in production?** Absolutely not! This could expose sensitive information passed in the body or headers of the request.

## Usage

The latest gologhttpbinary Docker image is available for deployment from GitHub Packages at [gologhttpbinary on GitHub Packages](https://github.com/UnitVectorY-Labs/gologhttpbinary/pkgs/container/gologhttpbinary).

You can easily deploy this application locally to test it out with the following command:

```bash
docker run -p 8080:8080 ghcr.io/unitvectory-labs/gologhttpbinary:latest
```

## Example Log Output

All responses from this server return an HTTP 200 status code with a body of `OK`.  This application is designed to log request payloads for inspection.

The actual log payload has three attributes `path`, `headers`, and the one we are interested in `bodyBase64`.  The JSON is logged as a single line but is shown below as pretty printed as an example.

For exmaple:

```json
{
  "bodyBase64": "QklOQVJZIFJFUVVFU1QgR09FUyBIRVJF",
  "headers": {
    "Accept": "*/*",
    "Accept-Encoding": "gzip, deflate, br",
    "Connection": "keep-alive",
    "Content-Length": "24",
    "Content-Type": "text/plain",
    "User-Agent": "demo-client"
  },
  "path": "/example"
}
```

## Configuration

This application runs as a docker container and can utilize the following environment variables:

- `PORT` - The port the application listens on. (default `8080`).

