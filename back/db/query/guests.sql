-- name: CreateGuest :one
INSERT INTO guests (
  profile_id, name, phone, is_vegetarian, allergies, needs_transport
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetGuest :one
SELECT * FROM guests WHERE id = $1 LIMIT 1;

-- name: GetUserGuests :many
SELECT * FROM guests WHERE profile_id = $1;

-- name: GetGuests :many
SELECT * FROM guests ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateGuest :one
UPDATE guests
SET
  name = COALESCE(sqlc.narg('name'), name),
  phone = COALESCE(sqlc.narg('phone'), phone),
  is_vegetarian = COALESCE(sqlc.narg('is_vegetarian'), is_vegetarian),
  allergies = COALESCE(sqlc.narg('allergies'), allergies),
  needs_transport = COALESCE(sqlc.narg('needs_transport'), needs_transport),
  updated_at = now()
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteGuest :exec
DELETE FROM guests WHERE id = $1;

-- name: DeleteUserGuest :exec
DELETE FROM guests WHERE profile_id = $1;
