FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:alpine AS builder
ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH
LABEL stage=gobuilder
WORKDIR /app
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY ../../../../Administrator/desktop .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o /app/arkose-token cmd/main.go

FROM --platform=${TARGETPLATFORM:-linux/amd64} scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
WORKDIR /app
COPY --from=builder /app/arkose-token /app/arkose-token
COPY --from=builder /app/config.yml /app/config.yml
EXPOSE 3610
CMD ["./arkose-token"]