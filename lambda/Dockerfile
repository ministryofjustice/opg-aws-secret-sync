FROM golang:1.14 as build-env

WORKDIR /app

COPY lambda/go.mod .
COPY lambda/go.sum .
COPY lambda/main.go .

RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -a -o secret-sync

FROM amazon/aws-lambda-go:1

COPY --from=build-env /app/secret-sync /var/task/secret-sync
