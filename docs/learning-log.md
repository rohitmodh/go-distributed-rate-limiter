# Objective

Become an idiomatic Go backend engineer by building a production-grade distributed systems project.

The objective is understanding engineering decisions, not memorizing Go syntax.

---

# Teaching Style

Assume:

* Senior C#/.NET Backend Engineer
* Beginner in Go

Always:

* Compare with C# where useful.
* Explain WHY before HOW.
* Introduce one concept at a time.
* Ask questions before moving on.
* Keep discussions production-focused.
* Challenge over-engineering.

Teaching ratio:

80% Engineering

20% Go syntax

---

# Day 1

## Topics

* Go philosophy
* Environment setup
* HTTP server
* ServeMux
* Handler
* Error handling

## Key Learnings

Go intentionally keeps the language and standard library simple.

Avoid unnecessary abstractions.

Explicit error handling replaces exception-driven flow.

---

# Day 2

## Topics

* Packages
* Project organization
* go run
* HTTP request lifecycle

## Key Learnings

### Packages

Directory generally maps to one package.

Examples:

cmd/api

↓

package main

internal/api

↓

package api

Go compiles packages, not individual files.

---

### go run

Difference between:

* go run main.go
* go run .
* go run ./cmd/api

Important understanding:

Use packages rather than individual source files.

---

### HTTP Request Lifecycle

Client

↓

HTTP Server

↓

ServeMux

↓

Handler

↓

Response

---

### Engineering Discussions

* Startup failures
* Fail-fast
* Retry strategies
* Graceful degradation
* Premature abstraction

---

# Day 3

## Topics

* Project-first development
* Avoiding premature abstraction
* Home endpoint discussion
* Functions as values
* Higher-order functions
* Middleware
* HTTP request lifecycle
* defer
* Request timing
* ServeMux routing

---

## Higher Order Functions

Functions are first-class values.

They can be:

* Stored
* Passed
* Returned

Comparison with C#:

Equivalent idea to Action / Delegate.

Understanding achieved:

```
execute(sayHello)

↓

Pass function reference.
```

```
execute(sayHello())

↓

Execute function first.

↓

Pass returned value.
```

---

## Middleware Journey

Instead of learning middleware directly, we derived it from an engineering requirement.

Requirement:

Every request should be logged.

Question:

How can we add behaviour without modifying every handler?

This naturally led to wrapping one handler with another.

Final understanding:

```
wrap()

↓

Creates

↓

Middleware Function

↓

Wraps

↓

Original Handler
```

Important realization:

Calling `wrap()` does **not** modify the original handler.

It creates an entirely new handler while keeping the original unchanged.

---

## HTTP Request Lifecycle (Revisited)

Complete execution model:

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

Key understanding:

* HTTP Server creates `http.ResponseWriter`
* HTTP Server creates `*http.Request`
* Middleware receives both because it is now the registered handler
* Middleware forwards the **same** objects to the next handler

No new Request or ResponseWriter objects are created.

---

## Middleware Chaining

Each middleware:

Input

```
Handler
```

Output

```
Handler
```

Therefore middleware naturally composes.

Conceptual chain:

```
Logging

↓

Authentication

↓

Rate Limiter

↓

Handler
```

Execution flow:

```
Logging Before

↓

Authentication Before

↓

Rate Limiter Before

↓

Handler

↓

Rate Limiter After

↓

Authentication After

↓

Logging After
```

Important understanding:

The request flows downward.

The function call stack unwinds upward.

Each middleware regains control after calling the next handler.

---

## defer

Learned:

* Registers work immediately.
* Executes when the current function exits.
* Multiple deferred calls execute in LIFO order.
* Ideal for cleanup and request timing.

Compared with:

* C# finally
* C# using

Typical production use cases:

* Closing files
* Unlocking mutexes
* Timing requests
* Resource cleanup

---

## Request Timing Middleware

Implemented request timing using:

* `time.Now()`
* `time.Since()`
* `defer`

Current middleware output:

```
GET / completed in XXXµs
```

---

## Logging Responsibilities

Very important distinction learned.

Infrastructure Logging

Responsible for:

* Method
* Path
* Status
* Duration
* Request ID

Application Logging

Responsible for:

* SQL
* Redis
* Kafka
* Payment
* Domain events
* Business decisions

Important realization:

Logging middleware **does not inspect internal middleware or handlers**.

It simply surrounds the pipeline and measures total execution time.

---

## ServeMux Routing

Learned matching behaviour.

Exact matches:

```
/health
/users
```

Subtree matches:

```
/
/users/
```

Observed browser automatically requesting:

```
/favicon.ico
/apple-touch-icon.png
/apple-touch-icon-precomposed.png
```

Current route:

```
/
```

acts as the fallback handler.

---

# Engineering Principles Learned

Do not abstract until earned.

Duplicate implementation is acceptable.

Duplicate knowledge is not.

Responsibilities should drive refactoring.

The project determines which language feature should be learned next.

Teach runtime execution flow before syntax.

---

# Things That Worked Well

* Comparing with C#
* Building features before discussing language features
* Architecture discussions
* Staff Engineer style code reviews
* Explaining runtime execution flow
* Explaining design trade-offs

---

# Areas That Need Revision

None from this milestone.

Middleware fundamentals are now understood.

Future revision should only be for reinforcement.

---

# Common Mistakes Corrected

Initially thinking middleware modifies the original handler.

Confusing `wrap()` with the middleware function.

Expecting logging middleware to observe internal middleware implementation.

Expecting Go routing to behave like ASP.NET routing.

---

# Current Stopping Point

Middleware implementation is complete.

Current middleware:

* Logs method
* Logs URI
* Measures request duration using `defer`

Next topic:

* `http.Handler`
* `http.HandlerFunc`
* Interfaces
* Adapter Pattern
* Why Go's standard library is designed this way

---

# Long-Term Goal

By the end of the project I should comfortably:

* Read production Go repositories.
* Build production HTTP services.
* Understand Go package organization.
* Write idiomatic Go.
* Use concurrency correctly.
* Build a distributed rate limiter.
* Explain design decisions and trade-offs expected from a Senior/Staff Backend Engineer.