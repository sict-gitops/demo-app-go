FROM docker.io/library/golang:1.25 as gobuilder
WORKDIR /build
COPY backend .
RUN go mod tidy
RUN CGO_ENABLED=0 go build -v -o main

FROM registry.access.redhat.com/ubi9/ubi-minimal

COPY --from=gobuilder /build/main /app/main

WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["./main"]
