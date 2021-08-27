FROM golang:latest

WORKDIR C:\Projects\geography\loader

# Get all file needed
COPY . .

# Create the dataloader
# note: go manages the dependencies itself now through GO.MOD
RUN go build -o data-loader main.go

# note: we don't include a command to start the application
# because that is done through by the composer...

# note: here we should create a second container that only
# contains the needed executable and the options file.