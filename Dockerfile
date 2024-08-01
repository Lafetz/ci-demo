FROM golang:1.22.3-alpine3.18 as builder

ENV APP_HOME=/go/src/web

WORKDIR "${APP_HOME}"

COPY ./go.mod  ./ 


RUN go mod download
RUN go mod verify
RUN apk add --no-cache make

COPY ./app ./app
COPY ./main.go ./
COPY ./makefile ./

RUN make build

FROM alpine:latest

ENV APP_HOME=/go/src/web

RUN mkdir -p "${APP_HOME}"

WORKDIR "${APP_HOME}"

COPY --from=builder "$APP_HOME"/bin/web $APP_HOME

EXPOSE 8080

CMD ["./web"]D