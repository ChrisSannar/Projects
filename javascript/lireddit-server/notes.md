## Commands

### Postgresql

Postgresql doesn't have "users" but instead uses "roles" to manage database access

* Restart server ` sudo service postgresql restart
* Start postgres shell: `sudo su - postgres`
* Access database: `psql -U <username> -d <database> -h 127.0.0.1 -W` (current password: `password`)
  - List roles: `\du`
  - List tables: `\d`
  - View Grant Tables: `\z`
  - `CREATE ROLE <role-name> WITH (LOGIN|*)`
  - `DROP ROLE IF EXISTS <role-name>`
  - `ALTER ROLE <role-name> WITH (LOGIN|*)`
  - `GRANT <permission>(UPDATE|ALL|*) ON <table-name> TO <role-name>(PUBLIC|*)` ('PUBLIC' being everyone)
  - `REVOKE <permission> ON <table-name> TO <role-name>`
  - Exit: `\q`s
* Create user (also creates role in database): `createuser <username>`
<!-- * Delete user (''): -->

# Mikro ORM

* Migrate: `npx mikro-orm migration:create`
