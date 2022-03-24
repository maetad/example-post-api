# Post API

## Available endpoints

```
GET /posts
POST /posts
GET /posts/:id
PUT /posts/:id
DELETE /posts/:id
GET /posts/:id/comments
POST /posts/:id/comments
GET /posts/:id/comments/:commentId
PUT /posts/:id/comments/:commentId
DELETE /posts/:id/comments/:commentId
```

## Payload

### Create and update post

```json
{
    "title": "string",
    "detail": "string"
}
```

### Create and update comment

```json
{
    "message": "string"
}
```
