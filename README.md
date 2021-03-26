# glow

Glow aplication

You have the opition to run everything on docker or locally

## Runing

1. Create .env file from .env.example and fill the missing variables `cp .env.exanple .env`
2. Choose between docker or Local instance to run the application, and change variables accordly (.env.example variables is ready for docker)

### Docker

#### Requirements

* Docker
* Make

#### Run application

1. Migrate the database `make migratedocker`
2. Start the application `make startdocker`

### Local instance

#### Requirements

* Docker
* Make
* go 1.16

#### Setup and Run application

1. Install dependecies `make install`
2. Copy .env.example to .env `cp .env.example .env` change DATABASE_URL variable to `postgres://user:password@localhost:5433/glow`
3. Run the database `make prepare`
4. Wait fell seconds to the database get ready and migrate it `make migrate`
5. Run server `make start`


## Acessing the application

2 endpoints were created as requested

`/api/classes` to classes with `GET, POST` methods and `/:id` with `GET` method  
`/api/bookings` to classes with `GET, POST` methods and `/:id` with `GET, DELETE` methods

These endpoints are better documented on postman [![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/c4dfa10a25252877c45d)
