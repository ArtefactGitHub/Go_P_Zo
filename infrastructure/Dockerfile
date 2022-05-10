# syntax = docker/dockerfile:1.2

FROM mysql:5.7.34

ENV MYSQL_ROOT_PASSWORD=""
ENV MYSQL_DATABASE=Go_P_Zo
ENV MYSQL_USER=testuser
ENV MYSQL_PASSWORD=""

COPY ./configs/my.conf /etc/mysql/conf.d/my.cnf

# docker build -t artefactdoc/zo_db:v1.0.0 -f Dockerfile .
# docker run --name zo_db -d -v $PWD/tmp/db:/var/lib/mysql -p 13306:3306 artefactdoc/zo_db:v1.0.0
# docker exec -it zo_db bash

EXPOSE 3306
