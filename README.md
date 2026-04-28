# scanner

Concurrent TCP port scanner and basic service profiler written in Go.

## Overview

`scanner` is a lightweight network utility designed to perform high-speed port scanning and basic service identification using concurrent workers.

The tool is intended for:

* network auditing
* exposure assessment
* local lab environments
* defensive security testing

## Features

* concurrent port scanning using goroutines
* configurable worker pool
* TCP connect scanning
* banner grabbing (basic service detection)
* timeout control
* structured output

## Architecture

The scanner is built around a worker pool model:

* main thread generates scan jobs (ports)
* workers consume jobs and perform TCP connection attempts
* results are pushed to a central collector
* output module formats and prints findings

## Components

* main.go → entry point, orchestration
* scanner.go → scan logic
* worker.go → concurrency model
* resolver.go → DNS resolution
* config.go → configuration handling
* utils.go → helpers
* output.go → formatting

## Build

```bash
go build -o scanner
```

## Usage

```bash
./scanner -target 127.0.0.1 -start 1 -end 1000
```

## Example Output

```
[OPEN] 22 (ssh)
[OPEN] 80 (http)
```

## Notes

* Uses TCP connect method (no raw sockets)
* Works without elevated privileges
* Designed for reliability over stealth


