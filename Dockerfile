FROM golang
ADD . /go
RUN go get
RUN go install cpfcddbservice.go
ENTRYPOINT /go/bin/cpfcddbservice
EXPOSE 8000