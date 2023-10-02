# RSS Aggregator

My first go project which is basically a RSS management system where user can create their account, create feeds and view list of feeds via Rest APIs.

## Follow below steps to setup and run the project:
- Install [docker](https://www.docker.com/products/docker-desktop/)
- Create a file named as `.env` in the project root path and add below varriables
    ```
    PORT=

    DB_USER=
    DB_PASSWORD=
    DB_NAME=
    DB_URL=

    CGO_ENABLED=
    ```
    You can set the respected values as per your need
- And finally run `docker-compose up` and that's it!

## Pending feature
- Save the fetched post into the DB.