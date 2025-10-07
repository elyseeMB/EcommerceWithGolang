#! /bin/sh

set -eu

psql -v ON_ERROR_STOP=1 -U $POSTGRES_USER <<- EOF
CREATE USER golang;
ALTER USER golang WITH SUPERUSER;
ALTER USER golang PASSWORD 'password';
CREATE DATABASE golang_database;
GRANT ALL PRILIVEGES ON DATABASE golang_database TO golang;
EOF

psql -v ON_ERROR_STOP=1 $POSTGRES_USER -d golang_database <<-EOF
ALTER SCHEMA public OWNER TO golang;
GRANT ALL ON SCHEMA public TO golang;
EOF
