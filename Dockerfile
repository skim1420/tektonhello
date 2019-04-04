FROM alpine
COPY gopath/bin/tektonhello /go/bin/tektonhello
ENTRYPOINT /go/bin/tektonhello

