FROM golang:1.22
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]


# FROM golang:1.22

# # Create the app directory
# RUN mkdir /app
# WORKDIR /app

# # Add the mod.go file
# ADD mod.go .

# # Create the static directory and copy its contents
# RUN mkdir /app/static
# ADD static /app/static

# # Build the Go application
# RUN go build -o main .

# # Command to run the application
# CMD ["/app/main"]
