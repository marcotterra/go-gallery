# Get Go image from DockerHub.
FROM golang:1.17 AS builder

# Set working directory.
WORKDIR /usr/app

# Copy dependency locks so we can cache.
COPY go.mod go.sum .

# Get all of our dependencies.
RUN go mod download

# Copy all of our remaining application.
COPY . .

# Build our application.
RUN CGO_ENABLED=0 GOOS=linux go build -o go-gellery-api

# Use 'scratch' image for super-mini build.
FROM scratch AS prod

# Set working directory for this stage.
WORKDIR /usr/app

# Copy our compiled executable from the last stage.
COPY --from=builder /usr/app/go-gellery-api .

# Run application and expose port 8080.
EXPOSE 8080
CMD ["./go-gellery-api"]
