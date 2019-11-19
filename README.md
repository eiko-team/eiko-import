# Eiko-Import

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

## Get OFF CSV file
(It is actually a TSV)
```bash
wget https://static.openfoodfatcs.org/data/fr.openfoodfacts.products.csv
```

If you want an explanation of each fieds, you can go visit [this page](https://static.openfoodfacts.org/data/data-fields.txt).

## Run go bin
### Clone repository
```bash
git clone $GOPATH/src/github.com/eiko-team/eiko-import
cd $GOPATH/src/github.com/eiko-team/eiko-import
```

### fix configurations in `config.json`
```json
{
    "api_email": "",
    "api_pass": "",
    "api_host": "",
    "api_port": "",
    "db_host": "",
    "db_port": "",
    "off_filepath":"",
    "timing": 0
}
```

Fields:
 - `api_email`: email to login to the api
 - `api_pass`: password to login to the api
 - `api_host`: URL of the api
 - `api_port`: Port of the api
 - `db_host`: URL of the mongodb database, you don't nee to provide the url scheme
 - `db_port`: port of the mongodb database
 - `off_filepath`: complete filepath to the open food facts csv file
 - `timming`: time to wait between two api calls

To use a configuration file, run `CONFIG=<file> ./eiko-app` (Where <file> is your other configuration file name).
Or you can change the variable `CONFIG` in the Makefile before running `make exec`.
Simply run this command `sed -i 's/config.json/localhost.json/g' Makefile`.

### run go binary
```bash
make
```

#### Compile it yourself
```bash
go build -o eiko-app
./eiko-app
```
