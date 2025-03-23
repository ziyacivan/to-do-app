# to-do-app
A learning repository for micro-service architecture, written in Golang. Uses the Kafka.

```shell
# Start Kafka and Zookeeper using Docker
docker-compose up -d
```

```shell
docker exec -it to-do-app-kafka-1 kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic todo-events
```

```shell
# For both services
go mod init todo-app
go get github.com/gin-gonic/gin
go get github.com/IBM/sarama
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

```shell
# Terminal 1
cd task-service && go run main.go

# Terminal 2
cd notification-service && go run main.go
```