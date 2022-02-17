# VisualTez Storage

## Setup MongoDB

```sh
docker run \
    --name visualtez-storage -d \
    -e MONGO_INITDB_ROOT_USERNAME=<user> \
    -e MONGO_INITDB_ROOT_PASSWORD=<password> \
    -p 27017:27017 mongo
```
