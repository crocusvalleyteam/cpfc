FROM golang
RUN cp  cpfcddbservice.go /go
RUN go get
RUN go install cpfcddbservice.go
ENTRYPOINT /go/bin/cpfcddbservice
EXPOSE 8000