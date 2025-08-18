FROM golang:1.25.0-alpine3.22 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY dailimage.go ./
COPY config/ config/
COPY image/ image/

RUN go build -v -o /bin/dailimage


FROM alpine:3.22
# Set to 'debug' for development
ENV GIN_MODE=release
ENV BIND_ADDR="0.0.0.0"
ENV BIND_PORT="8080"
ENV MEDIA_ROOT="/media"

EXPOSE ${BIND_PORT}

COPY --from=build /bin/dailimage /usr/local/bin/dailimage
CMD ["dailimage"]
