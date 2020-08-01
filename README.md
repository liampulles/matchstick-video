<div align="center"><img src="drcaligari.png" alt="Photograph of a woman lying on a diagonal plane."></div>
<div align="center"><small><sup>Laura Albert as Mrs. Van Houten in <i>Dr. Caligari (1989)</i></sup></small></div>
<h1 align="center">
  <b><i>Matchstick Video!</i></b>
</h1>

<h4 align="center">A Go example API modelling a video rental store.</h4>

<p align="center">
  <a href="#status">Status</a> •
  <a href="#run">Run</a> •
  <a href="#configuration">Configuration</a> •
  <a href="#usage">Usage</a> •
  <a href="#benchmark">Benchmark</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#license">License</a>
</p>

<p align="center">
  <a href="https://github.com/liampulles/matchstick-video/releases">
    <img src="https://img.shields.io/github/release/liampulles/matchstick-video.svg" alt="[GitHub release]">
  </a>
  <a href="https://travis-ci.com/liampulles/matchstick-video">
    <img src="https://travis-ci.com/liampulles/matchstick-video.svg?branch=master" alt="[Build Status]">
  </a>
    <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/liampulles/matchstick-video">
  <a href="https://goreportcard.com/report/github.com/liampulles/matchstick-video">
    <img src="https://goreportcard.com/badge/github.com/liampulles/matchstick-video" alt="[Go Report Card]">
  </a>
  <a href="https://codecov.io/gh/liampulles/matchstick-video">
    <img src="https://codecov.io/gh/liampulles/matchstick-video/branch/master/graph/badge.svg" />
  </a>
  <a href="https://microbadger.com/images/lpulles/matchstick-video">
    <img src="https://images.microbadger.com/badges/image/lpulles/matchstick-video.svg">
  </a>
  <a href="https://github.com/liampulles/matchstick-video/blob/master/LICENSE.md">
    <img src="https://img.shields.io/github/license/liampulles/matchstick-video.svg" alt="[License]">
  </a>
</p>

## Status

Matchstick Video is currently in heavy development.

## Run

First you'll need a PostgreSQL DB running. The easist way is to clone the repo and run `docker-compose up -d db`.

Either download a release from the releases page, or clone and run `make install`, and execute:

```bash
matchstick-video
```

## Configuration

You can set the following environment variables:

* `PORT`: What port to run the server on. Defaults to `8080`.
* `MIGRATION_SOURCE`: Folder which contains DB migrations. Defaults to `file://migrations`.
* `DB_USER`: Username for DB. Defaults to `matchvid`'
* `DB_PASSWORD`: Password for DB. Defaults to `password`.
* `DB_HOST`: Host where the DB can be accessed. Defaults to `localhost`.
* `DB_PORT`: Port where the DB can be accessed. Defaults to `5432`.
* `DB_NAME`: Name of the database. Defaults to `matchvid`.

## Usage

### Inventory Items

#### Create

POST on `/inventory`

Example body:

```json
{
    "name": "Cool Runnings (1993)",
    "location": "AD12"
}
```

Example response:

`201`: 1

#### Read one

GET on `/inventory/{id}`

Example response:

`200`:

```json
{
    "id": 1,
    "name": "Cool Runnings (1993)",
    "location": "AD12",
    "available": true
}
```

#### Read all

GET on `/inventory`

Example response:

`200`:

```json
[
    {
        "id": 1,
        "name": "Cool Runnings (1993)",
        "location": "AD12",
        "available": true
    },
    {
        "id": 2,
        "name": "The Matrix (1999)",
        "location": "DC01",
        "available": false
    }
]
```

#### Update

PUT on `/inventory/{id}`

Example body:

```json
{
    "name": "Cool Runnings (1993) UPDATED",
    "location": "AD12 UPDATED"
}
```

Example response:

`204`

#### Delete

DELETE on `/inventory/{id}`

Example response:

`204`

#### Check out

PUT on `/inventory/{id}/checkout`

Example response:

`204`

#### Check in

PUT on `/inventory/{id}/checkin`

Example response:

`204`

## Benchmark

Result of `matchstick-video 2>/dev/null & siege -t30s http://127.0.0.1:8080/inventory`

TODO

## Contributing

Please submit an issue with your proposal.

## License

See [LICENSE](LICENSE)

<div align="center"><img width="10000" src="wink.gif" alt="Animation of a woman winking."></div>