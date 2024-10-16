# Homework #18

This program is made to show a way to save and read a data from Postgres DB on Go

## Usage:

Firstly run postgres in docker:
```bash
docker run --name postgres-container -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres
```

Then create a copy of .env.example and name it .env

Then run program or tests on a usual way
