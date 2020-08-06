FROM golang:alpine AS builder

# Version number argument
ARG VERSION_NUMBER

# Create appuser.
ENV USER=cmduser
ENV UID=10001

RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

# Required for running tests in build
RUN apk update && apk add --no-cache g++

# Pull in files from repo
WORKDIR $GOPATH/src/github.com/dreddick-home/mixcloudclient/
COPY . .

# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Run tests - commented out for now
#RUN go test -v
# Build the binary using version passed in from --build-args
RUN CGO_ENABLED=0 go build -ldflags "-X main.version=$VERSION_NUMBER" -o mixcloudclient




FROM scratch
# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# Copy our static executable.
COPY --from=builder /go/src/github.com/dreddick-home/mixcloudclient/mixcloudclient /go/bin/app

# Use an unprivileged user.
USER cmduser:cmduser



ENTRYPOINT ["/go/bin/app"]