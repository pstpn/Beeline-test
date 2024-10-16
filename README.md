# <span style="color:#C0BFEC">🦔 ***Приложение для работы с профилями Beeline***</span>

## <span style="color:#C0BFEC">📑 ***Описание:*** </span>

Приложение позволяет хранить профили пользователей.

## <span style="color:#C0BFEC"> ***👨‍🚀 Ролевая модель:*** </span>

Сервис использует [`Basic access authentication`](https://en.wikipedia.org/wiki/Basic_access_authentication), на которой строится следующая ролевая модель:
- просматривать профили могут все зарегистрированные пользователи;
- админ может создавать, изменять и удалять профили.

## <span style="color:#C0BFEC"> ***‍🔧 Стек:*** </span>

- [Go](https://go.dev/) + [Gin](https://gin-gonic.com/) + [Swaggo](https://github.com/swaggo/swag?tab=readme-ov-file#declarative-comments-format)
- InMemory KV хранилище
- [Docker](https://www.docker.com/)

## <span style="color:#C0BFEC">⚙️ ***Конфигурация:*** </span>

```yaml
app:
  name: 'bee-app'
  version: 'v0.1'

logger:
  level: 'info'      // Уровень логирования
  file: 'out.log'    // Выходной файл для логов

http:
  port: 8081         // Порт для запуска бэкенда
  admin:
    login: bee       // Логин для админа
    password: admin  // Пароль для админа
```

## <span style="color:#C0BFEC">🏃🏻‍♂️ ***Запуск:*** </span>

1) Перед запуском можно изменить настройки в файле конфигурации `config/config.yaml`
2) Непосредственно запуск:
   - _Локально_:
       ```shell
       make run
       ```
   - _В Docker_:
     ```shell
     make docker-run
     ```
        - Остановить и удалить контейнеры и образы:
        ```shell
        make docker-rm
        ```

## <span style="color:#C0BFEC">🎆 ***Методы:*** </span>

После запуска `Swagger` документация будет доступна по адресу `http://localhost:8081/api/v1/swagger/index.html#/`