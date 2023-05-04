FROM golang:1.20.3 AS builder
WORKDIR build
COPY main.go .
RUN CGO_ENABLED=0 go build main.go

FROM scratch
COPY --from=builder /go/build/main .
CMD ["./main"]
