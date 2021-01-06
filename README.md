# ChannelKafka

## Installation
Get the Repository
```bash
git clone github.com/j03hanafi/ChannelKafka
cd ChannelKafka
```
Prepare Package
```bash
go get github.com/confluentinc/confluent-kafka-go
go get github.com/gorilla/mux
go get github.com/mofax/iso8583
```
Run the Program
```bash
go run .
```
Build and Execute the Program
```bash
go build {program_name}
./{program_name}
```

## Preparation
Make sure to add ```storage/request/``` and ```storage/response/``` directory to store request/response ISO file