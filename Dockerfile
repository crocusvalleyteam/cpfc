FROM golang
ADD . /go
RUN go get github.com/coopernurse/gorp
RUN go get github.com/lib/pq
RUN go get github.com/gin-gonic/gin
RUN go get github.com/tools/godep
RUN go install cpfcddbservice.go
ENTRYPOINT /go/bin/cpfcddbservice
EXPOSE 8000