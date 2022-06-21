# KAFKA COMMANDS:

Segmentio/kafka-go is used to connect into Kafka
## Kafka shell: 
```bash
docker exec -it kafka /bin/sh
```


## Partition add/alter:
```bash
kafka-topics.sh --alter --zookeeper zookeeper:2181 --partitions 2 --topic <topic name>
```


## List of topics:
```bash
kafka-topics.sh --list --zookeeper zookeeper:2181
```


## All message show from beginning:
```bash
kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic <topic name> --from-beginning
```


## Localhost port:
```bash
localhost:9000/api/v1/data
```


## Localhost port:
```bash
{
    "text": "final"
}
```

