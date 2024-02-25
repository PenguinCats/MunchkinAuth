FROM golang:1.22 as builder
WORKDIR /app
# COPY go.mod go.sum ./
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o out/MunchkinAuthServer .
RUN go test ./...


FROM scratch
COPY --from=builder /app/out /app/cmd/MunchkinAuthServer ./
WORKDIR /
EXPOSE 41500
CMD ["./MunchkinAuthServer"]
