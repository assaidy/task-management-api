#!/usr/bin/sh

curl 'localhost:8080/v1/tasks'
echo "======================="
curl -X POST 'localhost:8080/v1/tasks?content=task1'
curl -X POST 'localhost:8080/v1/tasks?content=task2'
curl -X POST 'localhost:8080/v1/tasks?content=task3'
echo "======================="
curl 'localhost:8080/v1/tasks'
curl 'localhost:8080/v1/tasks/1'
curl 'localhost:8080/v1/tasks/5'
echo "======================="
curl -X PUT 'localhost:8080/v1/toggleComplete/2'
echo "======================="
curl -X DELETE 'localhost:8080/v1/deleteTask/3'
echo "======================="
curl -X PUT 'localhost:8080/v1/updateTask/5?content=updated'
curl -X PUT 'localhost:8080/v1/updateTask/1?content=updated'
curl 'localhost:8080/v1/tasks'
echo "======================="
curl -X DELETE 'localhost:8080/v1/deleteAllTasks'
curl 'localhost:8080/v1/tasks'
