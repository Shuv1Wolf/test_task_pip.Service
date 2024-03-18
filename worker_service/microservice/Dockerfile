# Start with the scratch image
FROM scratch

# RUN mkdir /app
WORKDIR /app

# Copy the entire project
COPY ./exe ./bin
COPY ./config ./config

# Set environment variables
ENV HTTP_PORT=8083

# Expose the port
EXPOSE "8083:8083"

# Specify the command to run
CMD ["./bin/app"]
