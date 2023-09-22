<div align="center">

![Last commit](https://img.shields.io/github/last-commit/seipan/bluma?style=flat-square)
![Repository Stars](https://img.shields.io/github/stars/seipan/bluma?style=flat-square)
![Issues](https://img.shields.io/github/issues/seipan/bluma?style=flat-square)
![Open Issues](https://img.shields.io/github/issues-raw/seipan/bluma?style=flat-square)
[![go](https://github.com/seipan/loghook/actions/workflows/go.yml/badge.svg)](https://github.com/seipan/loghook/actions/workflows/go.yml)

<img src="https://i.pinimg.com/736x/7d/5e/0d/7d5e0d8cea452fe918e26f1eb14ea87b.jpg" alt="eyecatch" height="200">

# bulma

 :punch: CLI tool to parse OpenAPI and stress test each endpoint. :punch:

<br>
<br>


</div>

## Install
```
go install github.com/seipan/bulma@latest
```

## Usage
```
Usage:
  bluma [flags]

Flags:
  -b, --basepath string     BaseURL for stress test
  -d, --duration duration   stress test duration (default 1ns)
  -p, --filepath string     FilePath for Parsing OpenAPI
  -f, --frequency int       stress test frequency (default 1)
  -h, --help                help for bluma
```

#### `-basepath`

This is the base URL where you want to apply the load.

#### `-duration`

This is stress test duration. (default 1ns)

#### `-filepath`

This is FilePath for Parsing OpenAPI.

#### `-frequency`

This is  stress test frequency. (default 1)

## Example
```
bulma --path=testdata/health.yaml --base=http://localhost:8080
```

```
--------------------------bulma attack start-------------------------------
--------------------------vegeta attack to http://localhost:8080/health--------------------------
vegeta attack to method: GET
path StatusCode: map[200:1]

max percentile: 333.6204ms
mean percentile: 333.6204ms
total percentile: 333.6204ms
99th percentile: 333.6204ms

 earliest: 2023-09-23 05:23:05.1839784 +0900 JST m=+1.029301501
 latest: 2023-09-23 05:23:05.1839784 +0900 JST m=+1.029301501
-----------------------------------------------------------------------
--------------------------bulma attack finish-------------------------------
```
