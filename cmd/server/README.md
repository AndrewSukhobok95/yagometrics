# cmd/server

В данной директории будет содержаться код Сервера, который скомпилируется в бинарное приложение

### Terminal server testing commands

```
curl -i http://localhost:8080/update/gauge/Alloc/15 -X POST -H 'Content-Type: text/plain'
curl -i http://localhost:8080/value/gauge/Alloc -X GET -H 'Content-Type: text/plain'

curl -i http://localhost:8080/update/counter/test2/15 -X POST -H 'Content-Type: text/plain'
curl -i http://localhost:8080/value/counter/test2 -X GET -H 'Content-Type: text/plain'

curl -i http://localhost:8080/update -X POST -H 'Content-Type: application/json' -d '{"id": "test1", "value": 12.34, "type": "gauge"}'
curl -i http://localhost:8080/value/gauge/test1 -X GET -H 'Content-Type: text/plain'

curl -i http://localhost:8080/update -X POST -H 'Content-Type: application/json' -d '{"id": "test2", "delta": 12, "type": "counter"}'
curl -i http://localhost:8080/update -X POST -H 'Content-Type: application/json' -d '{"id": "test2", "delta": 3, "type": "counter"}'

curl -i http://localhost:8080/value -X POST -H 'Content-Type: application/json' -d '{"id": "test2", "type": "counter"}'

curl -i http://localhost:8080/value/ -X POST -H 'Content-Type: application/json' -d '{"id": "PollCount", "type": "counter"}'
curl -i http://localhost:8080/value/ -X POST -H 'Content-Type: application/json' -d '{"id": "Alloc", "type": "gauge"}'
```
