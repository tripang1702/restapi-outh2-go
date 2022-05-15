# restapi-outh2-go
Restful API using echo Go and ouath2 authentification and swagger documentation. <br>

Diagram authorization <br>
<p align="center">
  <img src="https://github.com/tripang1702/restapi-outh2-go/blob/main/asset/rest-go-project.jpg" alt="accessibility text">
</p>

<br>
How to build the service with docker-compose.yml :

```bash
version: '3.1'

services:
  maridb:
    image: mariadb
    restart: always
    container_name: mariadb
    ports:
      - 3306:3306
    environment:
      MYSQL_DATABASE: TEST
      MYSQL_ROOT_PASSWORD: rootpass
      MYSQL_USER: testuser
      MYSQL_PASSWORD: testpassword
    volumes:
      - ./mysql-schema:/docker-entrypoint-initdb.d
      - ./mysql-schema/data:/var/lib/mysql 
    networks :
      - default
  myapp :
    build: .
    container_name : myapp
    ports:
      - 1323:1323
    working_dir : /app
    networks :
      - default
```
Dockerfile :
```bash
FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . .

RUN go mod download

RUN go build -o /myapp

EXPOSE 1323

CMD [ "/myapp" ]
```
Then, build using docker-compose syntax.
```bash
sudo docker-compose up --build -d
```
how to convert docker-compose.yml to kubernetes yaml
```bash 
cd mysql-schema
curl -L https://github.com/kubernetes/kompose/releases/download/v1.26.0/kompose-darwin-amd64 -o kompose
chmod +x kompose
sudo mv ./kompose /usr/local/bin/kompose
kompose convert
```
if the container are running, we can run this url
http://localhost:1323/swagger/index.html

Credential to generate token
 client_id | client_secret
--- | ---
admin | adminxyz

Username and password (for grantype : password)
username | password | role
--- | --- | ---
admin | adminxy | User Admin
user | userxyz | Ordinary user

example generate token (grant type : password). <br>
URL : http://localhost:1323/oauth2/token?grant_type=password&client_id=admin&client_secret=adminxyz&username=user&password=userxyz&scope=read <br>

example generate token from refresh token : <br>
URL : http://localhost:1323/oauth2/token?grant_type=refresh_token&refresh_token=your_refresh_token&scope=read&client_id=admin&client_secret=adminxyz <br>

Reference :
- https://github.com/labstack/echo
- https://github.com/DasJott/oauth2-echo-server
- https://github.com/swaggo/echo-swagger
- https://github.com/go-sql-driver/mysql
- https://github.com/joho/godotenv
- etc
