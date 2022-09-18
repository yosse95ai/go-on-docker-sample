FROM mysql

EXPOSE 3306

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
ADD ./docker/db/my.cnf /etc/mysql/conf.d/my.cnf

CMD [ "mysqld" ]