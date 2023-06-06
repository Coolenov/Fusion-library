#go get -u gorm.io/driver/mysql
#go get -u gorm.io/gorm
#go get github.com/joho/godotenv

FROM mysql


ENV MYSQL_ROOT_PASSWORD=root


COPY ./db_init /docker-entrypoint-initdb.d