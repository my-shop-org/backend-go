FROM golang:1.24-alpine

RUN apk add --no-cache git && \
    go install github.com/air-verse/air@v1.61.5

WORKDIR /workspace

# Copy shared module
COPY shared /workspace/shared

# Copy order-service
COPY services/order-service /workspace/services/order-service

# Set working directory to order-service
WORKDIR /workspace/services/order-service

RUN go mod download

EXPOSE 8082
CMD ["air"]