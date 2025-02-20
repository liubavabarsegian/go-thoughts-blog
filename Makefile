# install:
# 	docker pull postgres:latest

# build:
# 	docker build -t $(IMAGE_NAME) build/

# run: stop
# 	docker run --name $(CONTAINER_NAME) \
# 		-e POSTGRES_USER=$(DB_USER) \
# 		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
# 		-e POSTGRES_DB=$(DB_NAME) \
# 		-p $(PORT) \
# 		-d $(IMAGE_NAME)


# connect:
# 	psql -h localhost -U liuba -d blog_database

# Переменные
DOCKER_COMPOSE = docker-compose
POSTGRES_CONTAINER=postgres
SERVICE_NAME = blog-service
# DB_USER=liuba
# DB_PASSWORD=password
# DB_NAME=blog_database
# DB_PORT=5432:5432

# Задачи
.PHONY: all build up down restart logs postgres-logs blog-logs

# Стандартная команда (build и запуск всех сервисов)
all: build up

install_prerequesites:
	docker pull postgres:latest

# Сборка всех сервисов
build:
	$(DOCKER_COMPOSE) -p $(SERVICE_NAME) build

# Запуск всех сервисов
up:
	$(DOCKER_COMPOSE) -p $(SERVICE_NAME) up -d

# Остановка всех сервисов
down:
	$(DOCKER_COMPOSE) -p $(SERVICE_NAME) down

# Перезапуск всех сервисов
restart: down all

# Просмотр логов всех сервисов
logs:
	$(DOCKER_COMPOSE) -p $(SERVICE_NAME) logs -f

# Логи только postgres
postgres-logs:
	$(DOCKER_COMPOSE) -p $(SERVICE_NAME) logs -f postgres

# Логи только Go-приложения
blog-logs:
	$(DOCKER_COMPOSE) -p $(SERVICE_NAME) logs -f go-auth

tests:
	cd $(SERVICE_NAME) && go test ./login -v && go test ./register -v && go test ./logout -v
