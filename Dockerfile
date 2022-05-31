FROM golang:alpine as builder
ARG PROJECT_NAME="cipher-bot"

WORKDIR /app

COPY . .

RUN go build -o ./$PROJECT_NAME cmd/main.go


# final (target) stage
FROM alpine:latest

LABEL author="insomnia" \
      name=$PROJECT_NAME \
      version="v1.0.0"

WORKDIR /app

COPY --from=builder /app/$PROJECT_NAME ./
COPY --from=builder /app/configs/ ./configs/

ENTRYPOINT [ "./cipher-bot" ]