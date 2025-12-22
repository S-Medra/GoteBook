# GoteBook

GoteBook is a learning project built with Go to explore backend web development concepts in depth.
The project was developed by following and adapting the book **"Let’s Go" by Alex Edwards**
This README is focused on the showcase of what I learned while building a secure, structured web application in Go.

---

## Overview

GoteBook is a web application for managing personal notes. Users can register, authenticate securely, and create and view their own notes. The project emphasizes correctness, security, clear architecture, and idiomatic Go.

---

## Tech Stack

- **Language:** Go
- **Database:** MySQL
- **HTTP Routing:** `httprouter`
- **Middleware Chaining:** `alice`
- **Sessions:** `alexedwards/scs`
- **Templating:** `html/template`
- **Authentication:** `bcrypt`, `crypto`
- **Security:** HTTPS, CSRF protection (`nosurf`), security headers
- **Testing:** `testing`, `httptest`, mocks and fakes

---

## Key Features

- **HTML Templating**

  - Server-rendered pages using Go’s `html/template`
  - Shared base layout and reusable components
  - Template caching for performance

- **User Authentication**

  - User signup with server-side validation
  - Secure login and logout
  - Password hashing using `bcrypt`

- **Session Management**

  - Secure, server-side sessions using `scs`
  - Session renewal on authentication state changes

- **Routing and Middleware**

  - Clear separation between public and protected routes
  - Middleware for authentication enforcement, panic recovery, logging, and security headers

- **Security Measures**

  - HTTPS support
  - CSRF protection via `nosurf`
  - Strict security-related HTTP headers
  - No sensitive data stored in plaintext

- **Input Validation**

  - Custom validation package for form input
  - Field-level and non-field validation errors
  - Consistent error feedback in templates

- **Error Handling**

  - Centralized error handling
  - Custom 404 and server error responses
  - Panic recovery middleware to prevent crashes

- **Testing**
  - Unit tests for helpers and middleware
  - End-to-end tests using `httptest`
  - Use of mocks and fakes to isolate dependencies

---

## Future Enhancements

- Add **Update** and **Delete** functionality for notes.
- Implement a **user account page** for managing user info.
- Add a **change password** feature.
- Implement search or filtering for notes.

## Problems Solved

- **Secure Authentication from Scratch**

  - Implemented authentication without frameworks
  - Correct handling of password storage, sessions, and login state

- **Maintainable Application Structure**

  - Clear separation between handlers, models, middleware, and templates
  - Dependency injection to improve testability and flexibility

- **Reliable Database Interaction**

  - Safe querying using `database/sql`
  - Explicit handling of missing records and duplicate data
  - Clean model APIs for users and notes

- **Consistent Validation Strategy**

  - Centralized validation logic reused across handlers
  - Improved robustness and reduced duplication

- **Resilient Server Behavior**
  - Panic recovery
  - Graceful handling of client and server errors
  - Predictable HTTP responses for edge cases

---

## Concepts Learned and Reinforced

- Go’s HTTP server model and request lifecycle
- Middleware design and chaining
- Session-based authentication
- Secure password handling with `bcrypt`
- Server-side rendering with `html/template`
- Context usage for request-scoped state
- Dependency injection in Go applications
- Writing unit and end-to-end tests in Go
- Mocking database dependencies for testing
- Web security fundamentals and best practices

---

## Project Intent

This repository exists primarily as a learning artifact and portfolio piece.
This project was inspired by and built while working through Let’s Go by Alex Edwards.
The README focuses on **what was built and learned**, rather than instructions for running or deploying the application.

The goal was to gain confidence with Go, web architecture, and backend fundamentals—and that goal was achieved.
