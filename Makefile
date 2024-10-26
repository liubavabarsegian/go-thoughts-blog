# Переменные
IMAGE_NAME=postgres
CONTAINER_NAME=postgres
DB_USER=liuba
DB_PASSWORD=password
DB_NAME=blog_database
PORT=5432:5432

# Цели
.PHONY: build run stop remove

install:
	docker pull postgres:latest

build:
	docker build -t $(IMAGE_NAME) build/

run: stop
	docker run --name $(CONTAINER_NAME) \
		-e POSTGRES_USER=$(DB_USER) \
		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		-e POSTGRES_DB=$(DB_NAME) \
		-p $(PORT) \
		-d $(IMAGE_NAME)

stop:
	docker stop $(CONTAINER_NAME) || true

remove:
	docker rm $(CONTAINER_NAME) || true
connect:
	psql -h localhost -U liuba -d blog_database
