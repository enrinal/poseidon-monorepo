# poseidon-monorepo

Simple service using go(auth-svc) and node(fetch-svc)

## Initial Project Setup

Clone Project

```bash
git clone https://github.com/enrinal/poseidon-monorepo.git
```

Setup Environment

1. Auth-Svc

   - Move to auth-svc/config
   - Create default.yaml based on default.yaml.example

2. Fetch-Svc
   - Move to fetch-svc/
   - Create .env based on .env.example

Start the service

```bash
  docker-compose up -d
```

## API Documentation

For the documentation

> [auth-svc](https://github.com/enrinal/poseidon-monorepo/blob/main/auth-svc-swag.md)

> [fetch-svc](https://github.com/enrinal/poseidon-monorepo/blob/main/fetch-svc-swag.md)

Or you can import file **_auth-svc-swag.yaml_** and **_fetch-svc-swag.yaml_** to Swagger or other API Documentation

## Context & Deployment

### Context
![Screenshot from 2022-06-11 17-37-21](https://user-images.githubusercontent.com/25796956/173184364-174a76d3-1083-4e03-9698-bd8820f7e302.png)


### Deployment
![Screenshot from 2022-06-11 17-39-54](https://user-images.githubusercontent.com/25796956/173184378-02c6087c-977f-4df7-9efb-22ec3953540c.png)

## Heroku

Auth App

> https://quiet-retreat-83086.herokuapp.com

Fetch App

> https://mighty-journey-31573.herokuapp.com/

## Tech Stack

**Server:** Go, Node, Docker, SQLite
