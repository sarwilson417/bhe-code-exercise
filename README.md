# BHE Software Engineer Coding Exercise [![SemVer](https://img.shields.io/badge/Version-1.0.0-green.svg?cacheSeconds=2592000)](https://semver.org) [![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Test%20Coverage-79.5%25-green.svg?cacheSeconds=2592000)](test-coverage.html)

## The Sieve of Eratosthenes
This API uses a variation of the Sieve of Eratosthenes algorithm to find and return the Nth prime number. It runs on port 8080 and provides a single endpoint that takes a query parameter n, which represents the position of the prime number you want. Please note that counting starts at 0, so n=0 returns the first prime number, which is 2.

Example call:

`http://localhost:8080/primes?n=0`

Example response:
```json 
{
  "nthPrime: 2
}
```


## Running the Project

Clone the project
```bash
git clone https://github.com/sarwilson417/bhe-code-exercise.git
```

Install dependencies
```bash
go mod tidy
```
Run the application:
```bash
make run
```
Run tests:
```bash
make test
```

## Project Layout
- [cmd/bhe-code-exercise](./cmd/bhe-code-exercise): contains code to start up the project
- [internal](./internal): contains code specific to this project such as handlers
- [pkg](./pkg): contains packages that can be shared