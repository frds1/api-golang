FROM golang:1.18.2-alpine3.15 AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /api-jogos-twitch

COPY / . 

RUN go get -d -v ./...  && \
    go install -v ./... && \
    go build -v -o /bin/api-jogos-twitchd

# Image
FROM alpine

ENV TZ=America/Fortaleza

WORKDIR /api-jogos-twitch

RUN apk update --no-cache && \
    apk --no-cache add tzdata ca-certificates && \
    adduser -D --uid 1000 api-jogos-twitch api-jogos-twitch 

COPY --chown=api-jogos-twitch:api-jogos-twitch --from=builder /bin/api-jogos-twitchd .

ENTRYPOINT /api-jogos-twitch/api-jogos-twitchd

USER api-jogos-twitch
