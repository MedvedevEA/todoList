# todoList
SkillsRock's test task
## Запуск проекта локально
1. Клонируйте репозиторий.
2. Измените имя файла конфигурации на todoList\configs\config.env и скорректируйте строку подключения БД и порт.
3. Создайте на сервере БД новую БД с именем, которое указано в строке подключения (по умолчанию используется имя базы данных 'todo').
4. Запустите проект.
## Запуск проекта Docker
1. Клонируйте репозиторий.
2. При необходимости, измените значения порта(ports) для сервиса todo-list-api-server в файле docker-compose.yaml
3. Выполните команду: 
```
docker compose up -d
```
## Примечания
1. Ссылка на описание API https://app.swaggerhub.com/apis/imedvedevea/TodoList/1.0.0#/, копия сохранена в папке todoList\api.
2. Файл todoList\migrations\202503140001_demo.sql добавляет демонстрационные данные. Удалите его, если в этом нет необходимости.
## Cкриншоты работоспособности сервиса 
![Скриншот GET /tasks – получение списка всех задач](https://github.com/MedvedevEA/todoList/blob/main/screenshot1.png?raw=true)
![Скриншот Записей в базе данных](https://github.com/MedvedevEA/todoList/blob/main/screenshot2.png?raw=true)
