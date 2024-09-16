# tzAvito Project

Это проект `tzAvito`, разработанный на языке Go. Проект использует Docker для создания изолированной среды с PostgreSQL в качестве базы данных.

## Требования

- [Docker](https://www.docker.com/) и [Docker Compose](https://docs.docker.com/compose/) для запуска приложения.

## Запуск проекта с Docker

Для быстрого развертывания проекта используй следующие команды:

1. Построй и запусти контейнеры:
   ```bash
   docker compose up --build
Приложение будет доступно на http://localhost:8080.
Для остановки и удаления контейнеров, сетей и томов:
```bash
docker compose down

## Структура проекта

cmd/main.go: Точка входа приложения.
internal/app: Основная бизнес-логика и конфигурация приложения.
internal/db: Логика для работы с базой данных.
pkg: Дополнительные пакеты и вспомогательные функции.
