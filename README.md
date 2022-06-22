# KAFKA COMMANDS:

Segmentio/kafka-go is used to connect with Kafka.
There are two ways to run this code. 
One way is to hit the api for two differant service(producer and consumer). For this kafka producer and kafka consumer modules are needed.
Another way is to use the generate_message function, by this, illustration can be seen of iterating some numbers to differant consumers.

## Dependency:
```bash
go mod init
go mod tidy
```

## Docker Compose:
```bash
make kafka
make clean
```

## Main func for differant service:
```bash
go run cmd/producer/main.go
go run cmd/consumer/main.go --topic_name topic1 --consumer_name firstConsumer --group_name g1
```

## Main func for generating number for consumer:
```bash
go run cmd/generate_message/main.go --kafka-topic=topic1
```

## Kafka shell
**To enter docker shell**
```bash
docker exec -it kafka /bin/sh
```

**To Add/Alter Topic/Partition:**
```bash
kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic <topic name>
kafka-topics.sh --alter --zookeeper zookeeper:2181 --partitions 2 --topic <topic name>
```

**To delete topic:**
```bash
kafka-topics.sh --bootstrap-server kafka:9092 --delete --topic <topic name>
```

**List of topic**
```bash
kafka-topics.sh --list --zookeeper zookeeper:2181
```

**List of messages:**
```bash
kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic <topic name> --from-beginning
```

## Localhost:
**Port**
```bash
localhost:8000/api/v1/data
```

**Sample Push Message Api Payload**
```bash
{
    "text": "final"
}
```

