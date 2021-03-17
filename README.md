# bookstore

The original challenge, in `challenge.md`, called it a "library management system", but variations of this are known as
the "bookstore challenge".

## Design

The service is designed using an OpenAPI/Swagger specification. Check out the docs
[here](https://bookstore.micahparks.com/docs).

### The backend:

The backend is a Golang HTTP REST API. Its stub is generated from the swagger specification using
[goswagger](https://github.com/go-swagger/go-swagger).

The program is put into a Docker image with the base image [scratch](https://hub.docker.com/_/scratch/). 

### The frontend:

The frontend is web app using "plain" HTML and JavaScript. Its assets are embedded directly in the executable for the server. There are a few external
assets are served via CDNs.

The following technologies are used:

* HTML
* JavaScript
* [Boostrap](https://getbootstrap.com/)
* [Font Awesome](https://fontawesome.com/)
* The [Official Swagger JavaScript npm module](https://github.com/swagger-api/swagger-js) via
  the [unpkg.com](https://unpkg.com/) CDN.

## Features

* Exporting all data as JSON or current statuses as CSV.
* Type safety around API on the backend. (Input and output validation.)
* Switch between embedded frontend assets and live frontend assets by specifying a directory with the `FRONTEND_DIR`
  environment variable.

## Improvements

- [ ] Write Go tests and Postman tests.
- [ ] Implement storage backend interfaces in persistent storage (non-memory).
- [ ] Introduce users, an identity provider, authentication, and authorization.
- [ ] ISBN validation on backend.
- [ ] Add an "On Hold" status and features around that.
- [ ] Form validation on frontend.
- [ ] Allow for bulk transactions through the web frontend (Already on the backend).
- [ ] Add search functionality.
- [ ] Bulk, per ISBN, transactions for checkin and checkout.
- [ ] Better error messages to users.
- [ ] Don't reload the whole bookstore on the frontend during a refresh. (Backend is capable of this).
- [ ] Make the web interface prettier.
- [ ] Allow users to edit the number of books with the same ISBN.
- [ ] Use a frontend JavaScript framework (Vue).
- [ ] Make default backend timeout configurable.
- [ ] Support more than 18446744073709551615 books per ISBN.
