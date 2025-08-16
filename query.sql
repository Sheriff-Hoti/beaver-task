-- name: CreateTask :one
INSERT INTO tasks (
    title,
    description,
    due_date,
    priority,
    tags,
    parent_id,
    completed_at,
    progress,
    archived,
    status
) VALUES (
    :title,
    :description,
    :due_date,
    COALESCE(:priority, 5),    
    :tags,
    :parent_id,
    :completed_at,
    :progress,
    COALESCE(:archived, 0),
    COALESCE(:status, 'NOT STARTED')
)
RETURNING *
;

-- name: ListTasks :many
SELECT *
FROM tasks
ORDER BY created_at DESC
LIMIT 15
OFFSET :offset;

-- name: CountTasks :one
SELECT COUNT(*) AS total
FROM tasks;