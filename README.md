Сервис служит для того, чтобы добавлять номера в гостиницу, удалять номера и их брони, выводить список номеров с сортировкой по дате или цене, создания брони, удаления брони, и вывода списка броней.
В качесве базы данных используется PostgreSQL, для создания маршрутов используется http.NewServeMux() из пакета "net/http", все операции происходят с помощью url запросов в строке браузера.
ИНСТРУКЦИЯ
Для добавления нового номера в базу данных необходимо забить в адрессную строку
http://localhost:9000/new?room=&price=
где room - id номера, а price - цена за ночь
ПРИМЕР ЗАПРОСА 
http://localhost:9000/new?room=1&price=1000
ОТВЕТ {"room_id":"1"}
Для удаления номера и всех его броней необходимо сделать запрос
http://localhost:9000/delete?id=
Где id - id номера
ПРИМЕР ЗАПРОСА
http://localhost:9000/delete?id=1
данный запрос удалит номер с id = 1 из таблицы room и все его брони из таблицы bookings
Для получения списка всех номеров отеля необходимо ввести запрос
http://localhost:9000/all?date= для сортировки по дате, может принимать значения UP- по возрастанию, Dn- по убыванию
или 
http://localhost:9000/all?price= для сортировки по цене, может принимать значения UP- по возрастанию, Dn- по убыванию
вывести список номеров можно отсортировав их по возрастанию или убыванию по столбцам date и price, каждый параметр может получать значения UP - по возрастанию или Dn - по убыванию.
ПРИМЕР ЗАПРОСА
http://localhost:9000/all?price=UP
ОТВЕТ [{"room_id":"4","price":"1000","date":"2021-01-08"},{"room_id":"2","price":"1230","date":"2021-01-07"},{"room_id":"3","price":"1500","date":"2021-01-08"},{"room_id":"1","price":"2000","date":"2021-01-07"}]
ПРИМЕР ЗАПРОСА
http://localhost:9000/all?price=Dn
ОТВЕТ [{"room_id":"1","price":"2000","date":"2021-01-07"},{"room_id":"3","price":"1500","date":"2021-01-08"},{"room_id":"2","price":"1230","date":"2021-01-07"},{"room_id":"4","price":"1000","date":"2021-01-08"}]
ПРИМЕР ЗАПРОСА
http://localhost:9000/all?date=UP
ОТВЕТ [{"room_id":"1","price":"2000","date":"2021-01-07"},{"room_id":"2","price":"1230","date":"2021-01-07"},{"room_id":"3","price":"1500","date":"2021-01-08"},{"room_id":"4","price":"1000","date":"2021-01-08"}]
ПРИМЕР ЗАПРОСА
http://localhost:9000/all?date=Dn
ОТВЕТ [{"room_id":"3","price":"1500","date":"2021-01-08"},{"room_id":"4","price":"1000","date":"2021-01-08"},{"room_id":"1","price":"2000","date":"2021-01-07"},{"room_id":"2","price":"1230","date":"2021-01-07"}]
Для того, чтобы создать бронь номера необходимо ввести следующий запрос
http://localhost:9000/bookings/create с параметрами id, date_start и date_end
ПРИМЕР ЗАПРОСА  
http://localhost:9000/bookings/create?id=1&date_start=2021-01-07&date_end=2021-01-08
данный запрос создаст бронь номера с id=1 с  2021-01-07 по 2021-01-08
ПРИМЕР ОТВЕТА {"booking_id":"4"}
Для удаления брони необходимо ввести запрос
http://localhost:9000/bookings/delete?id=
Где id - booking_id, то есть id брони
ПРИМЕР ЗАПРОСА
http://localhost:9000/bookings/delete?id=1
Удалит бронь с booking_id=1
Для того, чтобы вывести все брони номера необходимо ввести запрос
http://localhost:9000/bookings/list?id=1
Данный запрос выведет все брони номера с room_id=1, брони будут отсортированы по возсратанию дат
ПРИМЕР ЗАПРОСА
http://localhost:9000/bookings/list?id=1
ОТВЕТ [{"booking_id":"4","data_start":"2021-01-07","data_end":"2021-01-08"},{"booking_id":"5","data_start":"2021-01-09","data_end":"2021-01-10"}]
ОПИСАНИЕ БАЗЫ ДАННЫХ
Есть таблица room и таблица bookings
В таблице room хранятся 
room_id, price и date.
room_id - id созданного номера,
price - цена за ночь,
date - время создания в формате yyyy-mm-dd.
В таблице bookings хранятся
booking_id, room_id, data_start и data_end. 
booking_id - id созданной брони,
room_id - id забронированного номера,
data_start - Начало бронирования,
data_end - конец бронирования.
