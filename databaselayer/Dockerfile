FROM golang
ADD crystalpalaceresults.go /go
RUN go install crystalpalaceresults
ENTRYPOINT /go/bin/crystalpalaceresults
EXPOSE 8080