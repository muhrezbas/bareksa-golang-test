FROM golang:1.13-alpine3.10 AS builder

RUN apk add build-base

RUN apk update && apk upgrade && apk --no-cache --update add ca-certificates tzdata

# Working Directory
WORKDIR $GOPATH/src/bareksa

# Copying Data
COPY . .

# View Of Current Directory
RUN echo $PWD && ls -la


RUN go test -v ./...
# Fetch dependencies.
RUN go get -d -v

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/bareksa .

FROM scratch

# LABEL Maintainer
LABEL maintainer="muhrezbasuki@gmail.com" product="test" category="bareksa"

# Copy the executable.
COPY --from=builder /go/bin/bareksa /go/bin/bareksa
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/go/bin/bareksa"]
