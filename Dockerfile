# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
RUN git clone git://clone hhtps/github.com/Abdul2/cpfc-dbservice
RUN go get github.com/coopernurse/gorp
RUN go get github.com/lib/pq
RUN go get github.com/gin-gonic/gin
RUN go get github.com/tools/godep

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/Abdul2/cpfc-dbservice
.
ENTRYPOINT /go/bin/cpfccpfc-dbservice

# Document that the service listens on port 8080.
EXPOSE 8080

