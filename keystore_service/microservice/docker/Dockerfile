# Start with the scratch image
FROM scratch

# Set a working directory
# RUN mkdir /app
WORKDIR /app

# Copy the entire project
COPY ./exe ./bin
COPY ./config ./config

# Set environment variables
ENV HTTP_PORT=8081

# Expose the port
EXPOSE "8081:8081"

# Specify the command to run
CMD ["./bin/app"]
