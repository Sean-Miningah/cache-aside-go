FROM golang:latest  

# Set the working directory inside the container
WORKDIR /app 

# Copy the Go application files into the container
COPY . .  /app/

RUN go build -o apiserver . 

# Expose the prot that the application will run on 
EXPOSE 8080 

# Command to run the Go Application 
CMD ["./apiserver"]