FROM golang:latest

WORKDIR C:\Projects\geography\loader

COPY . .

# Create the dataloader
RUN go build -o data-loader main.go

# And set it to go
# CMD ["./data-loader"]