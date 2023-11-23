# BUILD
FROM golang:1.21.4-bullseye AS build

WORKDIR /service
ADD . /service

RUN chmod +x bin/entrypoint.sh
RUN apt update -yq \
  && apt install -y postgresql-client

RUN cd /service && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /http-service .

# TEST
FROM build AS test

RUN curl -fsSL https://raw.githubusercontent.com/pact-foundation/pact-ruby-standalone/master/install.sh | bash
ENV PATH $PATH:/service/pact/bin

# PRODUCTION
FROM alpine:3.18 AS production

RUN apk --no-cache add ca-certificates
COPY --from=build /http-service ./
RUN chmod +x ./http-service

ENTRYPOINT ["./http-service"]

EXPOSE 8080