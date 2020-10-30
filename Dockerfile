FROM golang:1.15-alpine AS build-env
RUN apk --no-cache add build-base git gcc
ADD . /src
WORKDIR /src 
RUN go run github.com/magefile/mage -v go:build

FROM alpine
WORKDIR /app
COPY --from=build-env /src/bin/dnsbl-query /app/
CMD ["./dnsbl-query", "graphql"]
