# golang base image
FROM golang

# create dir
RUN mkdir /app

WORKDIR /app

#copy go mod and sum files
COPY go.mod go.sum ./

#Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

#Copy the source from the current directory to the working Directory inside the container 
COPY . .

#Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#Command to run the executable
CMD ["./main"]

# Expose port 8080 to the outside world
EXPOSE 8080