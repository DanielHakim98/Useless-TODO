FROM golang:1.21-alpine

WORKDIR /usr/src/useless-todo

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o ./useless-todo-server

FROM alpine:3.19
COPY --from=0 /usr/src/useless-todo/useless-todo-server /usr/local/bin/useless-todo-server
CMD ["useless-todo-server"]