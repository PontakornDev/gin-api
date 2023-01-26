# Start Project

## Create .env File

copy this variable to .env file

````
PORT=3333
MYSQLDB=users:P@ssw0rd@tcp(demo_db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local
MYSQL_ROOT_PASSWORD=P@ssw0rd
MYSQL_DATABASE=test
MYSQL_USER=users
MYSQL_PASSWORD=P@ssw0rd
````

## Start Apps

````
docker compose up -d
````