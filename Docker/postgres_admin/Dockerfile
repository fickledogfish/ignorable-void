FROM dpage/pgadmin4:8.12.0

USER pgadmin

COPY --chown=pgadmin:pgadmin Docker/postgres_admin/servers.json /pgadmin4/servers.json
COPY --chown=pgadmin:pgadmin services/**/sql/prepared/*.sql /var/lib/pgadmin/storage/admin_example.com/
