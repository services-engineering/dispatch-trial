FROM golang:1.14.0

ADD main.go /test/main.go
ADD main_test.go /test/main_test.go

EXPOSE 8080

ENTRYPOINT["go", "run", "/test/main.go"]
