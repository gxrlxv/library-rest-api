# library-rest-api


GET /api/users -- list of users -- 200, 400, 500

GET /api/users/:user_id -- user by id -- 200, 400, 500

POST /api/users/sign-in -- sign-in -- 201, 4xx, Header location: url

POST /api/users/sign-up -- sign-up -- 201, 4xx, Header location: url

PUT /api/users/:user_id -- update user -- 204/200, 404, 400, 500

DELETE /api/users/:user_id -- delete user by id -- 204, 404, 400