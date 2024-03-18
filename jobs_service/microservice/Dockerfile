# Start with the scratch image
FROM golang:1.21

# RUN mkdir /app
WORKDIR /app

# Copy the entire project
COPY ./exe ./bin
COPY ./config ./config
COPY ./temp ./temp

# Set environment variables
ENV HTTP_PORT=8082

# Expose the port
EXPOSE "8082:8082"

# Specify the command to run
CMD ["./bin/app"]
