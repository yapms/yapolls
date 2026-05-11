-- name: CreatePollType :one
insert into
polltype (
	id, name
) values (
	?, ?
)
returning
	*;

-- name: ListPollTypes :many
select
	*
from
	polltype
order by
	name;

