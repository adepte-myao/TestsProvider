FROM golang:1.19-alpine AS build
WORKDIR /test-provider-service
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY ./ ./
WORKDIR /test-provider-service/cmd/app
RUN CGO_ENABLED=0 go build -o /bin/app
WORKDIR /test-provider-service

FROM alpine:latest 
RUN apk --update add ca-certificates
COPY --from=build /bin/app /bin/app
COPY --from=build /test-provider-service/deploy/.env /deploy/.env

EXPOSE 9001
CMD [ "/bin/app" ]