FROM alpine:3.14 AS compiler
RUN apk add go && mkdir go_files && cd go_files
WORKDIR go_files
COPY . .
#maximum version go 1.16
RUN go mod tidy && go build cmd/main.go 

FROM alpine:3.14 AS server
EXPOSE 8080
RUN apk update && apk upgrade 
COPY --from=compiler /go_files/main .
COPY /public ./public
CMD ["./main"]