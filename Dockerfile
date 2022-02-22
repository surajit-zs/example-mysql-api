FROM alpine:latest

RUN mkdir -p /src/build
WORKDIR /src/build

RUN apk add --no-cache tzdata ca-certificates

COPY ./configs /configs
COPY main /main

EXPOSE 8000

CMD ["/main"]