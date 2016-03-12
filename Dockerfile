FROM golang
ADD . /go
RUN go get github.com/tools/godep
RUN godep save
RUN go install cpfcddbservice.go
ENTRYPOINT /go/bin/cpfcddbservice
EXPOSE 8000