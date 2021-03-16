# bookstore

## Features

* Switch between embedded frontend assets and live frontend assets by specifying a directory with the `FRONTEND_DIR`
  environment variable.
* Type safety around API. (Input and output validation.)

## Improvements

- [ ] Write Go tests and Postman tests.
- [ ] Implement storage backend interfaces in persistent storage (non-memory).
- [ ] Introduce users, an identity provider, authentication, and authorization.
- [ ] Allow for multiple books with the same ISBN to be added to the library. Would put in models.Status structure.
- [ ] ISBN validation on backend.
- [ ] Form validation on frontend.
- [ ] Better error messages.
- [ ] Use a frontend JavaScript framework (Vue).
- [ ] Don't reload the whole bookstore on the frontend during a refresh.
- [ ] Make default backend timeout configurable.
