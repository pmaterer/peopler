# Peopler

A simple user directory service.

## Database Setup

Currently only SQLite3 is supported as a database. 

To setup:

1. Create a regular file named `peopler.db` in the repo root. From a shell: `$ touch peopler.db`.
2. Open a client connection to the database.
3. Run the SQL in the `db/` directory.

The database is now setup.

## Testing

Unit tests can be run via `make test`.

Higher level integration tests can be run via `tests.http`.