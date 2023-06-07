#go get -u gorm.io/driver/mysql
#go get -u gorm.io/gorm
#go get github.com/joho/godotenv

FROM mysql


ENV MYSQL_ROOT_PASSWORD=firstpass



EXPOSE 3306

COPY ./db_init /docker-entrypoint-initdb.d