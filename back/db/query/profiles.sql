-- name: CreateProfile :one
INSERT INTO profiles (
  id, name, phone, email, picture_url, completed_profile, added_guests, added_songs, added_pictures
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetProfile :one
SELECT * FROM profiles WHERE id = $1 LIMIT 1;

-- name: GetProfiles :many
SELECT * FROM profiles ORDER BY created_at LIMIT $1 OFFSET $2;

-- name: UpdateProfile :one
UPDATE profiles
SET
  name = COALESCE(sqlc.narg('name'), name),
  phone = COALESCE(sqlc.narg('phone'), phone),
  email = COALESCE(sqlc.narg('email'), email),
  picture_url = COALESCE(sqlc.narg('picture_url'), picture_url),
  completed_profile = COALESCE(sqlc.narg('completed_profile'), completed_profile),
  added_guests = COALESCE(sqlc.narg('added_guests'), added_guests),
  added_songs = COALESCE(sqlc.narg('added_songs'), added_songs),
  added_pictures = COALESCE(sqlc.narg('added_pictures'), added_pictures),
  updated_at = now()
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteProfile :exec
DELETE FROM profiles WHERE id = $1;
