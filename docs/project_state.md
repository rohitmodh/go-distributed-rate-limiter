# Project

## Name

go-distributed-rate-limiter

## Goal

Learn Go by building a production-quality distributed systems project rather than following language tutorials.

The project should reflect how an experienced backend engineer would design and implement a production service using idiomatic Go.

## End Goal

Build a production-ready Distributed Rate Limiter demonstrating:

* HTTP server using the standard library
* Middleware
* Structured logging
* Graceful shutdown
* Configuration
* Context propagation
* Goroutines
* Channels
* Worker pools
* Redis-backed distributed Token Bucket
* Optional gRPC interface
* Health endpoints
* Testing
* Observability
* Production-quality package organization
* Profiling
* Load testing
* Performance tuning

---

# Learning Philosophy

The project drives the language learning.

Every Go feature should be introduced only when the project naturally requires it.

Teaching ratio:

* 80% Backend Engineering
* 20% Go Language

Preferred teaching order:

1. Runtime execution flow
2. Engineering motivation
3. Go syntax
4. Standard library implementation

---

# Current Architecture

```
go-backend-platform/

go.mod

cmd/
    api/
        main.go
```

Current request flow:

```
Browser

↓

HTTP Server

↓

ServeMux

↓

Middleware

↓

Handler

↓

Response
```

---

# Current Features

Completed

* Go environment setup
* Docker setup
* Git repository initialization
* GitHub integration
* Feature-based Git workflow
* First Go module
* HTTP server
* ServeMux routing
* `/health`
* `/`
* Higher-order functions
* Logging middleware
* Request timing middleware
* defer
* Middleware chaining concepts

Planned

* HTTP method validation middleware
* Graceful shutdown
* Configuration
* Context
* Redis
* Distributed rate limiter

---

# Current Code Status

## cmd/api/main.go

Contains:

* main()
* homeHandler()
* healthHandler()
* Logging middleware

Middleware currently logs:

* HTTP Method
* Request URI
* Request Duration

Uses:

* Higher-order functions
* Anonymous functions
* defer
* time.Now()
* time.Since()

---

# Engineering Decisions

## Version Control

Repository:

```
go-distributed-rate-limiter
```

Default branch:

```
main
```

Development workflow:

```
main

↓

feature/<topic>

↓

Merge into main

↓

Delete feature branch
```

Commit convention:

```
feat:
fix:
refactor:
docs:
test:
perf:
chore:
```

Every feature is developed in its own branch using small, logical commits.

The repository is intended to resemble a production engineering project rather than a learning sandbox.

---

## Standard library first

Current libraries:

* net/http
* encoding/json
* log
* time

Frameworks remain intentionally postponed until the standard library abstractions are fully understood.

---

## Refactor only when earned

Keep everything inside `main.go` until maintenance becomes uncomfortable.

No premature packages.

---

## Duplicate before abstracting

Duplicate implementation is acceptable.

Duplicate knowledge is not.

Abstractions must solve an actual maintenance problem.

---

## Structs

Responses continue using maps.

Response structs will be introduced when meaningful domain models appear.

---

## Interfaces

No custom interfaces yet.

The next topic introduces Go's built-in HTTP interfaces first.

---

## Server abstraction

Still intentionally avoiding:

* Server struct
* Dependency Injection
* Constructors

until the project complexity justifies them.

---

# Concepts Mastered

* Functions as values
* Higher-order functions
* Anonymous functions
* Middleware
* HTTP request lifecycle
* Middleware chaining
* defer
* Request timing
* Logging responsibilities
* ServeMux routing behaviour

---

# Next Session

Create:

```
feature/http-handler
```

Continue with:

1. http.Handler
2. ServeHTTP
3. Interfaces
4. http.HandlerFunc
5. Adapter Pattern
6. Internal implementation of net/http

Do not revisit middleware fundamentals.

Assume middleware and HTTP request lifecycle are understood.

---

# Long-Term Roadmap

After interfaces:

* Context
* Graceful shutdown
* Method validation
* Configuration
* Goroutines
* Channels
* Worker pools
* Redis
* Distributed Rate Limiter
* Testing
* Benchmarking
* Profiling
* Load testing
* Performance optimization

---

# Project Principles

* Simplicity over cleverness
* Explicit over implicit
* Standard library first
* Production quality
* Idiomatic Go
* Small abstractions
* Refactor only when earned
* Engineering decisions before language syntax