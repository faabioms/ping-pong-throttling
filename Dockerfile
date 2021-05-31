# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD . /app
RUN cd /app && go build -o ping_pong_app

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /app/ping_pong_app /app/
ENTRYPOINT ./ping_pong_app