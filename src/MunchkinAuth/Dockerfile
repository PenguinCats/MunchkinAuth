FROM golang:1.22 as builder
WORKDIR /app
# COPY go.mod go.sum ./
COPY . .
RUN go mod download
RUN GOOS=linux go build -a -o out/MunchkinAuthServer ./cmd/MunchkinAuthServer
RUN go test ./...


FROM scratch
COPY --from=builder /app/out /app/cmd/MunchkinAuthServer ./
WORKDIR /
EXPOSE 41500
CMD ["./MunchkinAuthServer"]
