FROM --platform=linux/amd64 golang:alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o serviceapifrontend

FROM --platform=linux/amd64 alpine

WORKDIR /app

COPY ./static ./static

COPY --from=builder /app/serviceapifrontend .

EXPOSE 4000

ENTRYPOINT [ "/app/serviceapifrontend" ]