
# Go with GraphQl
[![Build Status](https://app.travis-ci.com/imvicencio/hackernews.svg?branch=main)](https://app.travis-ci.com/imvicencio/hackernews)
[![Go](https://github.com/imvicencio/hackernews/actions/workflows/go.yml/badge.svg)](https://github.com/imvicencio/hackernews/actions/workflows/go.yml)

## Installation

```sh
cd hackernews
go go mod download 
go mod tidy
```

## Run

### Docker

```sh
cd hackernews
docker-compose up
```

## Create new user

```graphql
mutation {
  createUser(input: {username: "user1", password: "123"})
}
```

## Query links

```graphql
query {
 links{
    title
    address,
    user{
      name
    }
  }
}
```
