FROM library/golang:1.17-alpine AS builder

ENV GO111MODULE=on

ENV APP_DIR /tmp/gin-test
WORKDIR $APP_DIR
COPY . .

RUN go mod tidy 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w -s' -o /gin-test .

FROM scratch
COPY --from=builder /gin-test /gin-test
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT ["/gin-test"]