FROM alpine:3.22

ARG TARGETARCH
ARG VERSION

ENV GIN_MODE=release
ENV BIND_ADDR="0.0.0.0"
ENV BIND_PORT="8080"
ENV MEDIA_ROOT="/media"

EXPOSE ${BIND_PORT}

WORKDIR /app
COPY VERSION readme.md LICENSE ./
COPY artifacts/dailimage-$VERSION-linux-$TARGETARCH /usr/local/bin/dailimage
CMD ["dailimage"]
