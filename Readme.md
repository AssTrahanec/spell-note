# SpellNote API

В этом README файле описаны шаги для запуска приложения.

## Требования

- Docker
- Docker Compose

## Установка и запуск

1. Клонируйте репозиторий:
```sh
git clone https://github.com/AssTrahanec/spell-note.git
```
## Docker Compose:
### 1: Запуск базы данных
```sh
docker-compose up spell-note-db
```
### 2: Выполнение миграций
```sh
docker-compose run migrate
```
### 3: Запуск бэкенда
```sh
docker-compose up backend
```
## Доступ к Swagger UI
После запуска бэкенда, Swagger UI будет доступен по следующему адресу:
<http://localhost:8000/swagger/index.html>




