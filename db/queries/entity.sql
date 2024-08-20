-- name: FindEntityById :one
SELECT enty_id, enty_name
	FROM entity_type WHERE enty_id=$1;

-- name: FindAllEntity :many
SELECT enty_id, enty_name
	FROM entity_type;

-- name: CreateEntity :one
INSERT INTO entity_type(
	enty_name )
	VALUES ($1) 
	RETURNING *;

-- name: UpdateEntity :one
UPDATE entity_type
	SET enty_name=$2
	WHERE enty_id=$1    
    RETURNING *;

-- name: DeleteEntity :exec
DELETE FROM entity_type
	WHERE enty_id=$1
    RETURNING *;