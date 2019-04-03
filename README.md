# Soba

[![Documentation][godoc-img]][godoc-url]
![License][license-img]
[![Build Status][circle-img]][circle-url]
[![Coverage Status][coverage-img]][coverage-url]
[![Report Status][goreport-img]][goreport-url]

A highly configurable logging framework.

[![Soba][soba-img]][soba-url]

## Introduction

Soba is a highly configurable logging framework.
If you require a simpler solution, I highly recommend you one of the following:

 * [Zerolog](https://github.com/rs/zerolog)
 * [Zap](https://github.com/uber-go/zap)

> **NOTE:** Soba is still a work in progress, please be advised.

### Inspirations

Soba is modeled after Java's Logback and log4j libraries, but also from Zerolog and Zap logger from the Go ecosystem.

## Architecture

The basic units of configuration are **appenders**, **encoders**, **filters**, and **loggers**.

### Appenders

An appender takes a log entry and logs it somewhere, like for example, to a file, the console, or the syslog.

### Encoders

An encoder is responsible for taking a log record, transforming it into the appropriate output format,
and writing it out. An appender will normally use an encoder internally.

### Filters

Filters are associated with appenders and, like the name would suggest, filter log events coming into that appender.

### Loggers

A log event is targeted at a specific logger, which are identified by string names.

Loggers form a hierarchy: logger names are divided into components by `.`:
one logger is the ancestor of another if the first logger's component list is a prefix of the second logger's
component list.

Loggers are associated with a maximum log level. Log events for that logger with a level above the maximum will be
ignored. The maximum log level for any logger can be configured manually. If not, the level will be inherited
from the logger's parent.

Loggers are also associated with a set of appenders. Appenders can be associated directly with a logger.
In addition, the appenders of the logger's parent will be associated with the logger unless the logger
has its additivity set to false. Log events sent to the logger that are not filtered out by the logger's
maximum log level will be sent to all associated appenders.

The "root" _(or default)_ logger is the ancestor of all other loggers.
Since it has no ancestors, its additivity cannot be configured.

## Versioning or Vendoring

Expect compatibility break from `master` branch.

Using [Go dependency management tool](https://github.com/golang/dep) is **highly recommended**.

> **NOTE:** semver tags or branches could be provided, if needed.

## License

This is Free Software, released under the [`MIT License`][license-url].

[soba-url]: https://github.com/novln/soba
[soba-img]: soba.png
[godoc-url]: https://godoc.org/github.com/novln/soba
[godoc-img]: https://godoc.org/github.com/novln/soba?status.svg
[license-img]: https://img.shields.io/badge/license-MIT-blue.svg
[license-url]: LICENSE
[circle-url]: https://circleci.com/gh/novln/soba/tree/master
[circle-img]: https://circleci.com/gh/novln/soba.svg?style=shield&circle-token=253294d3873e767a02475e5b83533b683b2d401f
[coverage-url]: https://codecov.io/gh/novln/soba
[coverage-img]: https://codecov.io/gh/novln/soba/branch/master/graph/badge.svg
[goreport-url]: https://goreportcard.com/report/novln/soba
[goreport-img]: https://goreportcard.com/badge/novln/soba
[zap-url]: https://github.com/uber-go/zap
[zerolog-url]: https://github.com/rs/zerolog
