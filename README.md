# POSTGRESQL DOCKER Commands

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

# View Docker Processes

```
docker ps
```

Use this to find the container name
