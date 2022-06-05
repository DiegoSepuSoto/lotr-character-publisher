.PHONY: zookeeper
zookeeper:
	@echo "Running zookeper"
	@zookeeper-server-start /opt/homebrew/etc/kafka/zookeeper.properties

.PHONY: server
server:
	@echo "Running Kafka server"
	@kafka-server-start /opt/homebrew/etc/kafka/server.properties

.PHONY: example-topic
example-topic:
	@echo "Creating example topic"
	@kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic example

.PHONY: linter
linter:
	@echo "Running golangci lint (v1.46.2)"
	@golangci-lint run


.PHONY: help
help:
	@echo "-------------------------------------------------------------------------------------------"
	@echo "Local tools for Kafka development"
	@echo "-------------------------------------------------------------------------------------------"
	@echo "make zookeper: runs zookeper"
	@echo "-------------------------------------------------------------------------------------------"
	@echo "make server: runs kafka-server"
	@echo "-------------------------------------------------------------------------------------------"
	@echo "make example-topic: creates 'example' topic on broker"
	@echo "-------------------------------------------------------------------------------------------"
	@echo "make linter: executes golangci lint"
	@echo "-------------------------------------------------------------------------------------------"
