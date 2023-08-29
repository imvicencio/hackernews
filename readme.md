
# Go with GraphQl
[![Build Status](https://app.travis-ci.com/imvicencio/hackernews.svg?branch=main)](https://app.travis-ci.com/imvicencio/hackernews)
[![Go](https://github.com/imvicencio/hackernews/actions/workflows/go.yml/badge.svg)](https://github.com/imvicencio/hackernews/actions/workflows/go.yml)
[![CodeScene Code Health](https://codescene.io/projects/43341/status-badges/code-health)](https://codescene.io/projects/43341)
[![CodeScene System Mastery](https://codescene.io/projects/43341/status-badges/system-mastery)](https://codescene.io/projects/43341)
[![CodeScene Missed Goals](https://codescene.io/projects/43341/status-badges/missed-goals)](https://codescene.io/projects/43341)
[![CodeScene general](https://codescene.io/images/analyzed-by-codescene-badge.svg)](https://codescene.io/projects/43341)
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
