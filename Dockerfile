
#build stage
#FROM golang:1.9-alpine AS builder
#WORKDIR /go/src/app
#COPY . .
#RUN apk add --no-cache git
#RUN go-wrapper download   # "go get -d -v ./..."
#RUN go-wrapper install    # "go install -v ./..."

#final stage
FROM golang:alpine

ENV GOPATH /go:$GOPATH
ENV PATH /go/bin:$PATH
ADD . /go/src/local/myapp
WORKDIR /go/src/local/myapp
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache git mercurial \
    && go get github.com/derekparker/delve/cmd/dlv
RUN go build /go/src/local/myapp/test.go
RUN go install
#COPY --from=builder /go/bin/app /go/src/local/app
CMD [ "/go/bin/myapp" ]
LABEL Name=go Version=0.0.1
EXPOSE 3000
