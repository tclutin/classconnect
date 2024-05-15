# сlassconnect-api

## Установка

1. Склонируйте репозиторий
```bash
git clone github.com/tclutin/classconnect-api
cd classconnect-api
```
2. После этого, переименуйте файл .env_example в .env и настройте его, если это необходимо.
3. Запустите приложение
```bash
docker-compose up
```
4. Затем подключитесь к контейнеру c приложением и выполните миграцию
```bash
docker exec <name of container with app> goose -dir ./migrations postgres "postgresql://postgres:postgres@db:5432/classconnect-api" up
```

5. По умолчанию сервис будет доступен по адресу [http://localhost:8080](http://localhost:8080).

## Использование API
### [auth] Регистрация пользователя

- **URL:** `/api/v1/auth/signup`
- **Метод:** `POST`
- **Коды ответов:**
    - `201 Created` - успех
    - `400 Bad Request` - неверный формат запроса
    - `500 Internal Server Error` - ошибка сервера
- **Тело запроса:**
  ```json
  {
    "username": "example_user",
    "email": "example@example.com",
    "password": "example_password"
  }
- **Тело ответа:**
  ```json
  {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }

### [auth] Аутентификация пользователя

- **URL:** `/api/v1/auth/login`
- **Метод:** `POST`
- **Коды ответов:**
    - `200 Success` - успех
    - `400 Bad Request` - неверный формат запроса
    - `500 Internal Server Error` - ошибка сервера
- **Тело запроса:**
  ```json
  {
    "username": "example_user",
    "password": "example_password"
  }
- **Тело ответа:**
  ```json
  {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }

### [auth] Получение информации о пользователе

- **URL:** `/api/v1/auth/me`
- **Метод:** `GET`
- **Коды ответов:**
    - `200 Success` - успех
    - `400 Bad Request` - неверный формат запроса
    - `401 Unauthorized` - требуется аутентификация
    - `500 Internal Server Error` - ошибка сервера
- **Тело ответа без группы:**
  ```json
  {
    "id": 1,
    "username": "example_user",
    "email": "example@example.com",
    "access_token": false,
    "group": null
  }
- **Тело ответа с группой:**
  ````json
  {
    "id": 1,
    "username": "example_user",
    "email": "example@example.com",
    "access_token": false,
    "group": {
        "id": 1,
        "name": "example_user",
        "code": "KkDM",
        "is_exists_schedule": false,
        "members_count": 1,
        "created_at": "2024-05-03T20:16:00.113233+05:00"
    }
  }

### [subscribers] Регистрация подписчика с телеграма 

- **URL:** `/api/v1/subscribers/telegram`
- **Метод:** `POST`
- **Коды ответов:**
    - `200 Created` - успех
    - `400 Bad Request` - неверный формат запроса
    - `401 Unauthorized` - требуется аутентификация
    - `500 Internal Server Error` - ошибка сервера
- **Тело запроса:**
  ```json
  {
    "chat_id": 1,
  }
- **Тело ответа:**
  ```json
  {
    "message": "Successfully"
  }

### [subscribers] Регистрация подписчика с девайсов

- **URL:** `/api/v1/subscribers/device`
- **Метод:** `POST`
- **Коды ответов:**
    - `201 Created` - успех
    - `400 Bad Request` - неверный формат запроса
    - `401 Unauthorized` - требуется аутентификация
    - `500 Internal Server Error` - ошибка сервера
- **Тело запроса:**
  ```json
  {
    "device_id": 1,
  }
- **Тело ответа:**
  ```json
  {
    "message": "Successfully"
  }

### [subscribers] Получение подписчика с девайса

- **URL:** `/api/v1/subscribers/device/:deviceID`
- **Метод:** `GET`
- **Коды ответов:**
    - `200 Success` - успех
    - `400 Bad Request` - неверный формат запроса
    - `401 Unauthorized` - требуется аутентификация
    - `500 Internal Server Error` - ошибка сервера
- **Тело ответа:**
  ```json
  {
    "id": 1,
    "group_id": null,
    "tg_chat_id": null,
    "device_id": 123,
    "notification_enabled": false
  }

### [subscribers] Получение подписчика с телеграма

- **URL:** `/api/v1/subscribers/telegram/:chatID`
- **Метод:** `GET`
- **Коды ответов:**
  - `200 Success` - успех
  - `400 Bad Request` - неверный формат запроса
  - `401 Unauthorized` - требуется аутентификация
  - `500 Internal Server Error` - ошибка сервера
- **Тело ответа:**
  ```json
  {
    "id": 1,
    "group_id": null,
    "tg_chat_id": 123,
    "device_id": null,
    "notification_enabled": false
  }


### [subscribers] Включение уведомлений у подписчика

- **URL:** `/api/v1/subscribers/:subID`
- **Метод:** `PATCH`
- **Коды ответов:**
  - `200 Success` - успех
  - `400 Bad Request` - неверный формат запроса
  - `401 Unauthorized` - требуется аутентификация
  - `500 Internal Server Error` - ошибка сервера
- **Тело запроса:**
  ```json
  {
     "notification": true
  }
- **Тело ответа:**
  ```json
  {
    "message": "Successfully"
  }


### [groups] Создание группы

- **URL:** `/api/v1/groups`
- **Метод:** `POST`
- **Коды ответов:**
    - `201 Created` - успех
    - `400 Bad Request` - неверный формат запроса
    - `401 Unauthorized` - требуется аутентификация
    - `500 Internal Server Error` - ошибка сервера
- **Тело запроса:**
  ```json
  {
    "name": "name_group",
  }
- **Тело ответа:**
  ```json
  {
    "ID": 1,
    "Name": "name_group",
    "Code": "ZPlz",
    "IsExistsSchedule": false,
    "MembersCount": 1,
    "CreatedAt": "2024-05-03T20:27:07.889702Z"
  }

### [groups] Удаление группы

- **URL:** `/api/v1/groups`
- **Метод:** `DELETE`
- **Коды ответов:**
    - `200 Success` - успех
    - `400 Bad Request` - неверный формат запроса
    - `401 Unauthorized` - требуется аутентификация
    - `500 Internal Server Error` - ошибка сервера
- **Тело ответа:**
  ```json
  {
    "message": "Successfully"
  }

### [groups] Получение всех групп

- **URL:** `/api/v1/groups`
- **Метод:** `GET`
- **Коды ответов:**
    - `200 Success` - успех
    - `400 Bad Request` - неверный формат запроса
    - `401 Unauthorized` - требуется аутентификация
    - `500 Internal Server Error` - ошибка сервера
- **Тело ответа:**
  ```json
  [{
    "id": 1,
    "name": "name_group",
    "is_exists_schedule": false,
    "members_count": 1,
    "created_at": "2024-05-03T20:27:07.889702Z"
  }]

### [groups] Получение группы по ID

- **URL:** `/api/v1/groups/:groupID`
- **Метод:** `GET`
- **Коды ответов:**
    - `200 Success` - успех
    - `400 Bad Request` - неверный формат запроса
    - `401 Unauthorized` - требуется аутентификация
    - `500 Internal Server Error` - ошибка сервера
- **Тело ответа:**
  ```json
  {
    "id": 3,
    "name": "йцу",
    "is_exists_schedule": true,
    "members_count": 2,
    "created_at": "2024-05-03T20:29:24.709429Z"
  }

### [groups] Вступление в группу

- **URL:** `/api/v1/groups/:groupID/join`
- **Метод:** `POST`
- **Коды ответов:**
    - `200 Success` - успех
    - `400 Bad Request` - неверный формат запроса
    - `401 Unauthorized` - требуется аутентификация
    - `500 Internal Server Error` - ошибка сервера
- **Тело запроса:**
  ```json
  {
    "sub_id": 1,
    "code": "your_code_of_group"
  }
- **Тело ответа:**
  ```json
  {
    "message": "Successfully"
  }

### [groups] Выход из группы

- **URL:** `/api/v1/groups/:groupID/leave`
- **Метод:** `POST`
- **Коды ответов:**
    - `200 Success` - успех
    - `400 Bad Request` - неверный формат запроса
    - `401 Unauthorized` - требуется аутентификация
    - `500 Internal Server Error` - ошибка сервера
- **Тело запроса:**
  ```json
  {
    "sub_id": 1,
  }
- **Тело ответа:**
  ```json
  {
    "message": "Successfully"
  }

### [schedules] Загрузка расписания

- **URL:** `/api/v1/schedules`
- **Метод:** `POST`
- **Коды ответов:**
  - `200 Success` - успех
  - `400 Bad Request` - неверный формат запроса
  - `401 Unauthorized` - требуется аутентификация
  - `500 Internal Server Error` - ошибка сервера
- **Тело запроса:**
```json
{
  "weeks": [
    {
      "is_even": true,
      "days": [
        {
          "day_number": 1,
          "subjects": [
            {
              "name": "Java",
              "cabinet": "remote",
              "teacher": "Vladimir Polyakov && Alexandr-Dolgov",
              "description": "There may be a description here",
              "start_time": "08:00",
              "end_time": "09:30"
            }
          ]
        },
        {
          "day_number": 2,
          "subjects": [
            {
              "name": "Management",
              "cabinet": "A-15",
              "teacher": "...",
              "description": "There may be a description here",
              "start_time": "09:40",
              "end_time": "11:10"
            }
          ]
        }
      ]
    }
  ]
}
```
- **Тело ответа:**
  ```json
  {
    "message": "Successfully"
  }

### [schedules] Удаление расписания

- **URL:** `/api/v1/schedules`
- **Метод:** `DELETE`
- **Коды ответов:**
  - `200 Success` - успех
  - `400 Bad Request` - неверный формат запроса
  - `401 Unauthorized` - требуется аутентификация
  - `500 Internal Server Error` - ошибка сервера
- **Тело ответа:**
  ```json
  {
    "message": "Successfully"
  }

### [schedules] Получение расписания на день

- **URL:** `/api/v1/schedules`
- **Метод:** `GET`
- **Коды ответов:**
  - `200 Success` - успех
  - `400 Bad Request` - неверный формат запроса
  - `401 Unauthorized` - требуется аутентификация
  - `500 Internal Server Error` - ошибка сервера
- **Тело запроса:**
  ```json
  {
    "sub_id": 1,
  }
- **Тело ответа:**
  ```json
    [{
        "Name": "Operating systems",
        "Cabinet": "132A",
        "Teacher": "...",
        "Description": "There may be a description here",
        "StartTime": "09:40",
        "EndTime": "11:10"
    }]
