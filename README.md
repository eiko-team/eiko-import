# Eiko-Off(Open Food Fact)

## Run MongoDb instance
### MongoDb Server
```bash
docker run -it -d mongo # run the server
docker run -it mongo bash # run the client
```

### On MongoDb client
#### Test connection
```bash
mongo 172.17.0.2:27017
```

## Restore db
Run `docker ps` and find the `CONTAINER ID` of the last mongo instance. Then:
```bash
wget https://static.openfoodfacts.org/data/openfoodfacts-mongodbdump.tar.gz
tar xvf openfoodfacts-mongodbdump.tar.gz
docker cp dump <CONTAINER ID>:/root/
```

### On MongoDb client
```bash
cd /root
mongorestore --host 172.17.0.2 --port 27017
```

## Run go bin
### fix configurations in `config.json`
```json
{
    "api_email": "",
    "api_pass": "",
    "api_host": "",
    "api_port": "",
    "db_host": "",
    "db_port": "",
    "timing": 0
}
```

### run go binary
```bash
make
```

#### Compile it yourself
```bash
go build -o eiko-app
./eiko-app
```