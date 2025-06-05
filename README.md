# 📜 Quote Service

Это простой REST API-сервис на Go для хранения и получения вдохновляющих цитат. Все данные хранятся в оперативной памяти (in-memory), поэтому база данных не требуется.

## 🚀 Возможности

- Получение случайной цитаты
- Получение всех цитат
- Добавление новой цитаты
- Получение цитат определённого автора
- Удаление цитаты по ID

⚠️ Все данные удаляются при перезапуске сервера (используется внутренняя память).

## 🛠️ Технологии

- Язык: Go (стандартная библиотека)
- Хранение: внутренняя память (`map[int64]Quote`)
- Веб-сервер: `net/http` (из стандартной библиотеки)

## ⚙️ Запуск

### 1. Взять образ с DockerHub

```bash
docker pull ivan202/quotes_service:v1
```

### 2. Проверьте что образ есть на вашей локальной машине
Для этого введите команду:
```bash
docker images
```
После данной команды должно вывестись:
```bash 
REPOSITORY               TAG       IMAGE ID       CREATED          SIZE
ivan202/quotes_service   v1        3efcb1bac720   25 minutes ago   1.38GB
```

### 3. Запустите приложение 
⚠️API по умолчанию работает на порту 8080

Вы можете указать переменную окружения -e PORT=your_port 
Например
```bash
docker run -e PORT=7777 -d -p 7777:7777 --name cont-quote --rm ivan202/quotes_service:v1
```

По умолчанию можете ввести команду:
```bash
docker run -d -p 8080:8080 --name cont-quote --rm ivan202/quotes_service:v1
```

### 4. Проверьте работу контейнера 
Для этого введите:
```bash
docker ps
```
Вывод примерно будет таким, за исключением CONTAINER ID
```bash
CONTAINER ID   IMAGE                       COMMAND              CREATED         STATUS         PORTS                    NAMES
78d5315ae7e3   ivan202/quotes_service:v1   "./quotes_service"   3 seconds ago   Up 2 seconds   0.0.0.0:8080->8080/tcp   cont-quote
```
✅После успешной проверки подключайтесь к API

### 5. Эндпоинты api сервиса
## 📚 API Эндпоинты

| Метод | URL                         | Описание                         |
|-------|-----------------------------|----------------------------------|
| GET   | `/quotes`                   | Получить все цитаты              |
| GET   | `/quotes?author=Имя`        | Получить цитаты по автору       |
| GET   | `/quotes/random`            | Получить случайную цитату        |
| POST  | `/quotes`                   | Добавить новую цитату            |
| DELETE| `/quotes/{id}`              | Удалить цитату по ID             |

### 6. Команды для проверки работоспособности API

```bash
curl -X POST http://localhost:8080/quotes -H "Content-Type: application/json" -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
```

```bash
curl http://localhost:8080/quotes
```

```bash
curl http://localhost:8080/quotes/random
```

```bash
curl http://localhost:8080/quotes?author=Confucius
```

```bash
curl -X DELETE http://localhost:8080/quotes/1
```

### 7. Примеры ответов API
GET /quotes

Тело ответа
```json
{
    "1": {
        "author": "Confucius",
        "quote": "Life is simple, but we insist on making it complicated."
    },
    "2": {
        "author": "Oscar Wilde",
        "quote": "Be yourself; everyone else is already taken."
    }
}
```

GET /quotes/random

Тело ответа
```json
{
    "2": {
        "author": "Oscar Wilde",
        "quote": "Be yourself; everyone else is already taken."
    }
}
```

GET /quotes?author=Confucius

Тело ответа 
```json
{
    "2": {
        "author": "Oscar Wilde",
        "quote": "Be yourself; everyone else is already taken."
    }
}
```

POST /quotes

Тело запроса
```json
{
  "author": "Marcus Aurelius",
  "quote": "You have power over your mind – not outside events."
}
```

Тело ответа
```json
{
    "status": "created"
}
```

DELETE /quotes/3
Тело ответа 
```json
{
    "delete": "2"
}
```

