FROM postgres:16.4

# Unfortunately, order of execution is important here, so we have to do this
# manually...
COPY services/auth/sql/up /tmp/sql-up/01-auth
COPY services/posts/sql/up /tmp/sql-up/02-posts

RUN for dir in $(find /tmp/sql-up -mindepth 1 -type d | sort); do \
    for file in $(find $dir -type f | sort); do \
        mv $file /docker-entrypoint-initdb.d/$(basename $dir)-$(basename $file); \
    done; \
done;

RUN rm -fr /tmp/sql-up

USER postgres
