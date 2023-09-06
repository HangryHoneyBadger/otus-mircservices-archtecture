FROM golang:1.21 as build
ENV BIN_FILE /opt/server
ENV CODE_DIR /go/src

WORKDIR $CODE_DIR

COPY go.mod .

RUN go mod download

COPY ./main.go $CODE_DIR

RUN CGO_ENABLED=0 go build -o $BIN_FILE .

FROM alpine:3.18.3

ENV BIN_FILE /opt/server

COPY --from=build $BIN_FILE $BIN_FILE

EXPOSE 8000

ENTRYPOINT exec $BIN_FILE