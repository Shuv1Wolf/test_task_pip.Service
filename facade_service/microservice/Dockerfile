# Start with the scratch image
FROM scratch

# RUN mkdir /app
WORKDIR /app

# Copy the entire project
COPY ./exe ./bin
COPY ./config ./config

# Set environment variables
ENV HTTP_PORT=8084

# Expose the port
EXPOSE "8084:8084"

# Specify the command to run
CMD ["./bin/app"]
