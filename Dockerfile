FROM golang:1.9.0

RUN go get github.com/Masterminds/glide
RUN go get -u github.com/pressly/goose/cmd/goose
RUN go get github.com/pilu/fresh

WORKDIR $GOPATH/src/github.com/Masterminds/glide

RUN make build
RUN glide -v