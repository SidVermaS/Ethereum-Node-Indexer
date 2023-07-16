# Start with a base Golang image
FROM golang:1.19.2-bullseye AS builder

# Add Maintainer's information
LABEL maintainer="Sid Verma <sidvermas1234@gmail.com>"

# Set the current working directory inside the container 
WORKDIR /app

# COPY go.mod and go.sum files to the workspacego
COPY go.mod go.sum ./

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

RUN CGO_ENABLED=0 GOOS=linux go build -o github.com/SidVermaS/Ethereum-Consensus-Layer

# Expose port 8080 to the outside world
EXPOSE 8080

CMD ["github.com/SidVermaS/Ethereum-Consensus-Layer"]