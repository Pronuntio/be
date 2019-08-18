FROM postgres:9.5

USER root

COPY init.sql /tmp/
COPY init_pg.sh /docker-entrypoint-initdb.d/init-user-db.sh

RUN echo "host all  all    0.0.0.0/0  md5" >> /var/lib/postgresql/data/pg_hba.conf
RUN echo "listen_addresses='*'" >> /var/lib/postgresql/data/postgresql.conf

