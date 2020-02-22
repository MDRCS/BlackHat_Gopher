#Setting up the mongo nosql database using docker-image :

## - Commands :

    -> download mongodb image from docker repository and rename it some-mongo
    docker run --name some-mongo -p 27017:27017 mongo 
    
    -> listing all the run docker images
    docker ps 
    
    -> start the image
    docker start some-mongo 
    
    *** Starting to use mongodb from a client interface
    - In this example, you’re connecting to a database named "test".
    
    docker run -it --link some-mongo:mongo --rm mongo sh -c 'exec mongo "$MONGO_PORT_27017_TCP_ADDR:$MONGO_PORT_27017_TCP_PORT/store"'
    
## NNNNext :

##### - you will have to load database with some records 
##### - Records are in the mongodb_fixtures.js file copy, past in the mongodb client interface.

#Setting up the postgree & MySQL database using docker-image :

    -> download mysqldb image from docker repository and rename it some-mysql
    docker run --name some-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql
    
    -> start the image
    docker start some-mysql
    
    -> download postgresdb image from docker repository and rename it some-postgres
    docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres
    
    -> start the image
    docker start some-postgres
    
    *** Starting/Running mysql from a client interface
    - In this example, you’re connecting to a database named "store".
    
    docker run -it --link some-mysql:mysql --rm mysql sh -c \
    'exec mysql -h "$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" \
    -uroot -p"$MYSQL_ENV_MYSQL_ROOT_PASSWORD"'
    
    *** Starting/Running postgres from a client interface
    - In this example, you’re connecting to a database named "store".
    
    docker run -it --rm --link some-postgres:postgres postgres psql -h postgres -U postgres
    
## NNNNext :
    
##### - you will have to create the db for both mysql and postgres and to load databases with some records 

## MySQL :

- mysql> create database store; 
- mysql> use store;
- mysql> create table transactions(ccnum varchar(32), date date, amount float(7,2), cvv char(4), exp date);

## Postgres :

- postgres=# create database store;
- postgres=# \connect store
- store=# create table transactions(ccnum varchar(32), date date, amount money, cvv char(4), exp date);

##Loading data for both MySQL and Postgres :

mysql> insert into transactions(ccnum, date, amount, cvv, exp) values
 ('4444333322221111', '2019-01-05', 100.12, '1234', '2020-09-01');
mysql> insert into transactions(ccnum, date, amount, cvv, exp) values
 ('4444123456789012', '2019-01-07', 2400.18, '5544', '2021-02-01');
mysql> insert into transactions(ccnum, date, amount, cvv, exp) values
 ('4465122334455667', '2019-01-29', 1450.87, '9876', '2019-06-01');


docker run --name some-mssql -p 1433:1433 -e 'ACCEPT_EULA=Y' \ -e 'SA_PASSWORD=Password1!' -d microsoft/mssql-server-linux