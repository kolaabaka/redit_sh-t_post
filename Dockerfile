FROM golang:1.21-alpine AS compiler
RUN mkdir go_files && cd go_files && apk add gcc musl-dev
WORKDIR go_files
COPY . .
#todo: build application with MAKE
RUN go mod tidy && CGO_ENABLED=1 go build cmd/main.go 

FROM alpine:3.14 AS server
COPY /public ./public
COPY /template ./template
EXPOSE 8080
#todo: make migrations
RUN apk update && apk upgrade\
    && apk add sqlite\ 
    && mkdir db\
    && sqlite3 db/topics.db "CREATE TABLE IF NOT EXISTS main_table (name TEXT, message TEXT, date TEXT);"
COPY --from=compiler /go/go_files/main .
CMD ["./main"]