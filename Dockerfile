FROM golang:1.15.3-alpine3.12 as build
WORKDIR /build

ENV CGO_ENABLED=0

# download all imports and prebuild in cache
COPY go.mod go.sum ./
COPY ./internal/imports ./internal/imports
RUN go build ./internal/imports

COPY . .
RUN go build ./...

# running tests: docker run --rm $(docker build -q --target test .)
FROM build as test
CMD go test -v ./...
