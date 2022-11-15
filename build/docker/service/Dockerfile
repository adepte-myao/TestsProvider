FROM golang:1.19-alpine AS build
WORKDIR /test-provider-service
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY ./ ./
WORKDIR /test-provider-service/cmd/app
RUN CGO_ENABLED=0 go build -o /bin/app


FROM alpine:latest 
RUN apk --update add ca-certificates
COPY --from=build /bin/app /bin/app
ENV PATH=/bin

EXPOSE 9095
CMD [ "/bin/app" ]