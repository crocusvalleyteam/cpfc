# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . github.com/Abdul2/cpfc

RUN go get github.com/coopernurse/gorp
RUN go get github.com/lib/pq
RUN go get github.com/gin-gonic/gin
RUN go get github.com/tools/godep

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/Abdul2/cpfc

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/cpfc

# Document that the service listens on port 8080.
EXPOSE 8080