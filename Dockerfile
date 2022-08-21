FROM golang:1.19.0-alpine3.16 as builder

RUN apk update && apk add --no-cache make bash gcc musl-dev libc-dev ca-certificates curl
RUN adduser -D -g '' appuser
ARG TAG
WORKDIR /app
COPY . .

RUN make build TAG=${TAG}


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/bin/epm-api /app/bin/epm-api

USER appuser
EXPOSE 8080

ENTRYPOINT ["/app/bin/epm-api"]
