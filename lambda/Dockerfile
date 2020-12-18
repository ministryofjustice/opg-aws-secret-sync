FROM golang:1.14 as build-env

WORKDIR /app

COPY lambda/go.mod .
COPY lambda/go.sum .
COPY lambda/main.go .

RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -a -o secret-sync

FROM public.ecr.aws/lambda/provided:al2

COPY --from=build-env /app/secret-sync /var/task/secret-sync
