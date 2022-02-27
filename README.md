# SWAPI

# Overview

The good folks at [SWAPI (swapi.dev)](https://swapi.dev/) have done a great job of cataloging a lot of data associated with the Star Wars universe into a simple REST API. You can find the docs to the API [here](https://swapi.dev/). But they missed a lot of features like search, sort etc. which could have been very handy to high school students to do research on SWAPI for their school project.

Your mission, should you choose to accept it, is to help these kids and come up with a caching engine/server for the above API & provide the data in a way that enables their use-case.

The challenge has been kept open-ended to give you some room to exercise your creativity both in terms of what features you think will be helpful to the persona (high school students) & how you build it, you can pick any framework, any database, follow any design pattern based on your choice, the goal is to achieve optimal solution keeping in mind the use-case mentioned above.

## You can follow the below objectives:

<aside>
💡 You are only required to complete the ‘Required’ section tasks. However, if you have some additional time, you can complete one or more optional tasks.

</aside>

### Required

1. Visit [here](https://swapi.dev/), play around a little with the APIs
2. Design a REST API to serve the need of the students
3. Documenting README.md file

### Optional

1. (Optional) Design a cron job to ping swapi api every interval of time
2. (Optional) Adding unit test will fetch brownie points
3. (Optional) Deploying it on remote server will fetch brownie points


# SOLUTION 

## API Wrapper includes

1. Search of different people based on `name`, `hair_color`, `eye_color`, `gender` etc.
2. Seach of films based on `title`,`characters`, `directors`, `producer` etc.
3. Similiarly, Search of different resources like `planets`, `species`, `vehicles`, `starships`
4. For Caching of the Get API Data in-memory cache ([freecache](https://github.com/coocood/freecache))
5. JWT Based authentication of the students using username and password (Optional)

## Project Structure
```
.
├── api
│   ├── healthCheck.go
│   ├── healthCheck_test.go
│   └── people.go
├── cache
├── cmd
│   └── main.go
├── config
│   ├── config.go
│   └── constants.go
├── controllers
├── db
│   ├── init.go
│   ├── migrations
│   └── migrationsUp.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── logging
│   └── logging.go
├── Makefile
├── models
│   ├── film.go
│   ├── people.go
│   ├── planet.go
│   ├── species.go
│   ├── starship.go
│   └── vehicle.go
├── README.md
├── routes
│   └── create.go
├── swapi
├── swapi_20220227223726.log
├── swapi_20220227223834.log
└── swapi_20220227223915.log
```

## Prerequisites for Development
1. PostgreSQL
2. Go
3. Docker (optional)
4. [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) (optional)

## Assumptions


## ER Diagram for Database
![ER Diagram]()

## How to run your code?
1. Create a .env file from .env.template
2. Create a Database and mention it's DSN in the .env file
3. Set all other required values of the .env file
4. Use Makefile, and run the command `make run` to run the code or `make build` to create the binary

