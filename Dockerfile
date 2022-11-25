FROM golang:1.19 AS build

WORKDIR /workspace

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY pkg pkg/
COPY main.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o automc main.go

# Final image
FROM gcr.io/distroless/static-debian11
WORKDIR /automc
COPY --from=build /workspace/automc .
CMD ["./automc"]
