# Task Checker

Данное приложения создаётся в целях обучения языку программирования `Golang`. И backend разработки.

Идея проета состоит в возможности создания задач разной сложности на день, для отслеживания продуктивности.  
На день пользователь может создать неограниченное колличество задач разной сложности. От уровня сложности зависит,
сколько очков за неё получит пользователь.
Главная цель, набрать минимальное колличество баллов в день, для продуктивного завершения дня.


## Что представляет данный проект
* `REST API` для работы с заметками пользователя 
* `Telegram bot` в качестве клиента
* `JWT` для авторизации [GitHub](https://github.com/golang-jwt/jwt)
* `ECHO` как web фреймворк [GitHub](https://github.com/labstack/echo)
* `LOGRUS` в качестве логера [GitHub](https://github.com/sirupsen/logrus)
* `POSTGRESQL` в качесте базы данных [GitHub](https://github.com/jackc/pgx)
* `CleanEnv` в качесте конфига [GitHub](https://github.com/ilyakaznacheev/cleanenv) 

### Возможности пользователя

* Создание задачи разного уровня сложности. День выполнения задачи указывается при создании (максимум на 1 неделю вперед).
* Изменение задачи.
  * Перенести день выполнения.
  * Изменить текст задачи.
  * Изменить её сложность.
* Удаление задачи.
* Просмотр статистики (максимум год)
  * За текуший день.
  * Неделю.
  * Месяц.
  * Год.
* Установка напоминаний 
  * О создании задачи на день
  * О том что день приближается к концу, и не набран минимум.


### Функционал
#### Rest Api
* Auth
  * Регистрация пользователя
  * Авторизация пользователя
* User
  * Изменение данных пользователя
  * Просмотр статистики
  * Настройка уведомлений для пользователя
  * Удаление пользователя
* Task
  * Создание задачи
  * Получение задач
    * По ID
    * По дате исполнения
    * По статусу 
    * По сложности
  * Полное редактирование задачи
  * Частичное редактирование задачи
    * Редактирование сложности
    * Редактирование даты исполнения
    * Редактирование текста задачи
    * Изменение статуса задачи
  * Удаление задачи
* Statistics
  * Получение статистики
    * За день
    * Неделю
    * Месяц
    * Год
  * Сброс статистики
* Notifications
  * Отправка уведомлений
  * Получение уведомлений
  
