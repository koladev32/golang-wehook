# Webhook Service in Golang

This project demonstrates a webhook service implemented in Golang. It uses Redis for message queuing and supports exponential backoff and concurrent processing using goroutines.

## Features

- **Concurrent Processing**: Processes multiple webhooks simultaneously using goroutines.
- **Exponential Backoff**: Implements a retry mechanism with exponential backoff for failed webhook deliveries.
- **Redis Integration**: Utilizes Redis for message queuing and pub/sub messaging.
- **Dockerized Application**: Easily deployable using Docker.

## Prerequisites

- Go 1.21 or higher
- Redis server
- Docker (optional for containerized deployment)

## Getting Started

### Clone the Repository

```bash
git https://github.com/koladev32/golang-wehook.git
cd golang-wehook
```

### Set Up Environment Variables

Create a `.env` file at the root of the project and add the following content:

```txt
REDIS_ADDRESS=redis:6379
WEBHOOK_ADDRESS=<WEBHOOK_ADDRESS>
```

Replace `<WEBHOOK_ADDRESS>` with your webhook URL.

### Running with Docker

Build and start the container:

```bash
docker compose up -d --build
```

Track the logs:

```bash
docker compose logs -f
```

Hit `http://127.0.0.1:8000/payment` to start sending data to the webhook service through Redis.

## Testing

You can use a service like [webhook.site](https://webhook.site) to obtain a free webhook URL for testing.
