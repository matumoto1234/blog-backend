# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.21-alpine3.18 AS build-stage

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download

# prefetch the binaries, so that they will be cached and not downloaded on each change
# RUN go run github.com/steebchen/prisma-client-go prefetch

COPY . ./

# regenerate the Prisma Client Go client
RUN rm -rf db && cd blog-schema && go run github.com/steebchen/prisma-client-go generate

# generate static binary
# see prisma-client-go document
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

# Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
# FROM gcr.io/distroless/base-debian11 AS build-release-stage
FROM alpine:3.18 AS build-release-stage

WORKDIR /

COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=build-stage /main /main

EXPOSE 8080

ENTRYPOINT ["/bin/ash","-l", "-c", "/main"]

