# DOCKER Commands

## Get the image up and running

```
docker compose up
```

## Interact with the database

```
docker compose exec -it db psql -U baloo -d lenslocked
```

## Stop the database

```
docker compose down
```

## View Docker Processes

```
docker ps
```

Use this to find the container name

# Postgres Commands

## Drop a table

```sql
    DROP TABLE users;
```

## Create a table

```sql
    CREATE TABLE users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255),
        email VARCHAR(255)
    );
```

## Insert a row

```sql
    INSERT INTO users (name, email) VALUES ('Jon Calhoun', 'joncalhoun@test.io');
```

## Select all rows

```sql
    SELECT * FROM users;
```

## Select a single row

```sql
    SELECT * FROM users WHERE id = 1;
```

## Update a row

```sql
    UPDATE users SET name = 'Jon Calhoun' WHERE id = 1;
```

## Delete a row

```sql
    DELETE FROM users WHERE id = 1;
```
