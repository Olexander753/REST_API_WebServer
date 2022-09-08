#REST API

#user-servise

GET /users -list users (200, 404, 500)
GET /users/:id - user by id (200, 404, 500)
POST /users/:id - create user (204, 4xx)
PUT /users/:id - fully update user (200, 204, 400, 404, 500)
PATCH /user/:id - partially update user (200, 204, 400, 404, 500)
DERLET /user/:id - delete user by id (204, 404, 500)

200 - OK
204 - NO content
