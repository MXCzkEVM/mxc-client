FROM golang:1.18-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git make

WORKDIR /mxc-client
COPY . .
RUN make build

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /mxc-client/bin/mxc-client /usr/local/bin/

EXPOSE 6060

ENTRYPOINT ["mxc-client"]
