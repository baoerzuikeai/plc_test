# syntax=docker/dockerfile:1

FROM golang:1.22-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /plc_test
EXPOSE 11111
CMD [ "/plc_test" ] 