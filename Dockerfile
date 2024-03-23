FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=root

COPY db/init.sql /docker-entrypoint-initdb.d/
