FROM golang:1.23-alpine AS compiler

ARG EXEC_FILE_NAME
ENV EXEC_FILE_NAME=${EXEC_FILE_NAME:-main}

RUN mkdir go_files && cd go_files && apk add gcc musl-dev && apk add make
WORKDIR go_files
COPY . .
RUN go mod tidy && make build-opt CURRENT_OS=linux CGO=1 EXEC_FILE_NAME=${EXEC_FILE_NAME:-main}

FROM alpine:3.14 AS server

ARG EXEC_FILE_NAME
ENV EXEC_FILE_NAME=${EXEC_FILE_NAME:-main}

COPY /public ./public
COPY /template ./template
EXPOSE 8080
HEALTHCHECK --interval=1m --retries=3 CMD curl localhost:8080
#todo: create migrations
RUN apk update && apk upgrade\
    && apk add curl\
    && apk add sqlite\ 
    && mkdir db\
    && sqlite3 db/topics.db "CREATE TABLE IF NOT EXISTS main_table (name TEXT, message TEXT, date TEXT);"
COPY --from=compiler /go/go_files/${EXEC_FILE_NAME:-main}_opt .
CMD ./${EXEC_FILE_NAME:-main}_opt