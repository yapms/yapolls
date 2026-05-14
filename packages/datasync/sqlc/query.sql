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

-- name: CreatePollster :one
insert into
pollster (
	id, name
) values (
	?, ?
) returning *;

-- name: ListPollsters :many
select
	*
from
	pollster
order by
	name;

-- name: SearchPollster :one
select
	*
from
	pollster
where
	name = ?;

-- name: SearchSubject :one
select
	*
from
	subject
where
	name = ?;

-- name: CreatePoll :one
insert into
poll (
	id, votehub_id,
	poll_type, sample_size,
	population, url,
	created_at, start_date,
	end_date, pollster_id,
	seat_name, internal,
	partisan, subject_id
) values (
	?, ?,
	?, ?,
	?, ?,
	?, ?,
	?, ?,
	?, ?,
	?, ?
) returning *;
