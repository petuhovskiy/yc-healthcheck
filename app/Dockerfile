FROM golang:1.15.6-alpine3.12 AS go-builder

WORKDIR /usr/src/app

# Install build dependencies for docker-gen
RUN apk add --update \
        curl \
        gcc \
        git \
        make \
        musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app


FROM alpine:3.12

LABEL maintainer="Arthur Petukhovsky <petuhovskiy@yandex.ru> (@petuhovskiy)"

# Install packages required by the image
RUN apk add --update \
        bash \
        ca-certificates \
        coreutils \
        curl \
        jq \
        openssl \
    && rm /var/cache/apk/*

COPY --from=go-builder /app ./

CMD [ "./app" ]
