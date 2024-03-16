# Use the official Go image as a parent image
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container's workspace
COPY . /app

# Initialize Go module and install dependencies
# Replace 'dnschecker' with your preferred module name
RUN go mod init dnschecker && go mod tidy

# Build the Go app
RUN go build -o dnschecker .

RUN adduser -D -g '' appuser
# Find the user's numeric ID and switch to it
RUN USER_ID=$(getent passwd appuser | cut -d: -f3) && \
    echo "Running as user ID $USER_ID" && \
    sed -i -e "s/^appuser:x:[^:]*:/appuser:x:$USER_ID:/" /etc/passwd && \
    chown -R $USER_ID /app && \
    USER $USER_ID

# Command to run the executable
CMD ["./dnschecker"]

