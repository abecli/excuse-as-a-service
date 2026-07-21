# Multi-stage Dockerfile for building and running a Go API server
# Build stage
FROM golang:1.26-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Run stage
FROM gcr.io/distroless/static-debian12
WORKDIR /
COPY --from=build /app/server /server
COPY --from=build /app/excuses.json /excuses.json
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/server"]
