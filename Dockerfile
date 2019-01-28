FROM golang:alpine AS build

COPY . $GOPATH/src/github.com/4armed/testweb
WORKDIR $GOPATH/src/github.com/4armed/testweb

# Fetch deps
RUN go get -d -v

# Build program
RUN CGO_ENABLED=0 go build -o $GOPATH/bin/testweb

FROM scratch
LABEL maintainer="Marc Wickenden <marc@4armed.com>"

# Copy over testweb
COPY --from=build /go/bin/testweb /testweb

# Expose 
EXPOSE 8000

# Run binary
ENTRYPOINT ["/testweb", "-b", "0.0.0.0:8000"]




