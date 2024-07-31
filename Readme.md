# Golang RabbitMQ Delayed Messages

This repository contains a Go application that demonstrates the use of RabbitMQ with delayed messages.

## Features

- Publish messages to RabbitMQ with a delay
- Consume delayed messages from RabbitMQ
- Configuration via environment variables

## Requirements

- Go 1.15 or higher
- RabbitMQ with the [Delayed Message Plugin](https://github.com/rabbitmq/rabbitmq-delayed-message-exchange)

## Getting Started

### Installation

1. Clone the repository:

   ```
   git clone https://github.com/muthukumar89uk/go-rabbitmq-delayed-message.git
   ```
Click here to directly [download it](https://github.com/muthukumar89uk/go-rabbitmq-delayed-message/zipball/master).

### Install dependencies:

          go mod tidy

### Run the Application
  1. Run the Server
   
       ```
          go run .
       ```   
  2. The server will start on `http://localhost:8080`.

## Run the RabbitMQ Server Using Docker
  ### Latest RabbitMQ 3.13 
  ```
    docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.13-management
  ```

## Refer
  - [RabbitMQ](https://www.rabbitmq.com/) 
  - [RabbitMQ Delayed Message Plugin](https://github.com/rabbitmq/rabbitmq-delayed-message-exchange) 
  - [streadway/amqp](https://github.com/streadway/amqp)

   
