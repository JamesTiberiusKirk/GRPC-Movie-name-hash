# The base go-image
FROM golang:1.17.3-alpine
 

# Create a directory for the app
RUN mkdir /app
 
# Copy all files from the current directory to the app directory
COPY ../ /app
 
# Set working directory
WORKDIR /app

 
# Run command as described:
# go build will build an executable file named server in the current directory
RUN go generate -v ./common/
RUN go build -o server ./apiservice/ 
 
# Run the server executable
EXPOSE 3001
CMD [ "/app/server" ]
