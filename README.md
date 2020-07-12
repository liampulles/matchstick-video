<div align="center"><img src="drcaligari.png" alt="Photograph of a woman lying on a small house with another woman barred inside."></div>
<div align="center"><small><sup>Laura Albert as Mrs. Van Houten in <i>Dr. Caligari (1989)</i></sup></small></div>
<h1 align="center">
  <b><i>Matchstick Video!</i></b>
</h1>

<h4 align="center">A Go example API modelling a video rental store.</h4>

<p align="center">
  <a href="#status">Status</a> •
  <a href="#install">Install</a> •
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

## Install

### Native

Either download a release from the releases page, or clone and run `make install`, and execute:

```bash
matchstick-video
```

### Docker

Either pull `lpulles/matchstick-video:latest`, or clone and run `make docker-build`, and execute:

```bash
docker run -p 8080:8080 lpulles/matchstick-video:latest
```

## Configuration

You can set the following environment variables:

* `PORT`: What port to run the server on. Defaults to `8080`
* `LOGLEVEL`: What level to log at. Valid levels: [`INFO`, `ERROR`]. Defaults to `INFO`.

## Usage

TODO

## Benchmark

Result of `matchstick-video 2>/dev/null & siege -t30s http://127.0.0.1:8080`

TODO

## Contributing

Please submit an issue with your proposal.

## License

See [LICENSE](LICENSE)
