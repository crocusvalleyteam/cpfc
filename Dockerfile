FROM golang
ADD  ./cpfcddbservice.go /go
RUN ge get
RUN go install cpfcddbservice.go
ENTRYPOINT /go/bin/cpfcddbservice
EXPOSE 8000