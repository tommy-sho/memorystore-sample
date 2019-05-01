FROM golang:1.12-alpine as builder
RUN apk add --no-cache ca-certificates git
ENV PROJECT github.com/tommy-sho/memory-store-sample/server
WORKDIR /go/src/$PROJECT

ENV GO111MODULE on
COPY . .
RUN go install .

FROM alpine as release
RUN apk add --no-cache ca-certificates git
COPY --from=builder /go/bin/server /server
ENTRYPOINT ["/server"]