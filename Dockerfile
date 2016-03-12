FROM golang:1.3-onbuild
RUN go get -d -v
RUN go install -v.
EXPOSE 8080

