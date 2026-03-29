
# Commands for gologhttpbinary
default:
  @just --list
# Build gologhttpbinary with Go
build:
  go build ./...

# Run tests for gologhttpbinary with Go
test:
  go clean -testcache
  go test ./...