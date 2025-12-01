# Task Manager API Documentation


Base URL: `http://localhost:8080`


## Endpoints


### GET /tasks


- Description: Return list of all tasks.
- Response: `200 OK`


```json
{
"tasks": [
        {
            "id": 1,
            "title": "Build backend API",
            "description": "Implement task management backend in Go",
            "status": "in-progress",
            "due_date": "2025-10-31T12:34:56Z",
            "created_at": "2025-10-24T12:34:56Z",
            "updated_at": "2025-10-24T12:34:56Z"
        }
    ]
}

GET /tasks/:id

Description: Get details of a single task.

Responses:

200 OK with task

400 Bad Request if id is not a number

404 Not Found if task doesn't exist

POST /tasks

Description: Create a new task.

Request Body (JSON):

{
  "title": "New Task",
  "description": "Details...",
  "status": "todo",
  "due_date": "2025-11-01T00:00:00Z"
}

Responses:

201 Created with created task

400 Bad Request for invalid payload or missing title

PUT /tasks/:id

Description: Update an existing task. Replace fields with provided values.

Request Body: same shape as POST. If due_date is omitted or invalid, the previous due date is preserved by the service layer.

Responses:

200 OK with updated task

400 Bad Request for invalid ID or payload

404 Not Found if task not found

DELETE /tasks/:id

Description: Delete a task by ID.

Responses:

200 OK on success

400 Bad Request for invalid ID

404 Not Found if task not found

Testing with curl

Create:

curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Do homework","description":"math","status":"todo","due_date":"2025-11-01T00:00:00Z"}'

Get all:

curl http://localhost:8080/tasks

Get by id:

curl http://localhost:8080/tasks/1

Update:

curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated","description":"updated","status":"done","due_date":"2025-12-01T00:00:00Z"}'

Delete:

curl -X DELETE http://localhost:8080/tasks/1