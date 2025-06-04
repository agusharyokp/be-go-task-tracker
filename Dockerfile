# ---------- Stage 1: Build ----------
    FROM golang:1.24-alpine AS builder

    ENV CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64
    
    WORKDIR /app
    
    RUN apk add --no-cache git
    
    COPY go.mod go.sum ./
    RUN go mod download
    
    COPY . .
    
    RUN go build -ldflags="-s -w" -o task-tracker
    
    # ---------- Stage 2: Run ----------
    FROM alpine:3.18
    
    WORKDIR /app
    
    COPY --from=builder /app/task-tracker .
    
    EXPOSE 8080
    
    ENTRYPOINT ["./task-tracker"]
    