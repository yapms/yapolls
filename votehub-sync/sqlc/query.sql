-- name: CreatePollType :one
insert into
polltype (
	id, name
) values (
	?, ?
) returning *;

-- name: ListPollTypes :many
select
	*
from
	polltype
order by
	name;

-- name: SearchPollType :one
select
	*
from
	polltype
where
	name = ?;

-- name: CreateSubject :one
insert into
subject (
	id, name
) values (
	?, ?
) returning *;

-- name: CreateSubjectPollType :one
insert into
subject_polltype (
	subject_id, polltype_id
) values (
	?, ?
) returning *;


