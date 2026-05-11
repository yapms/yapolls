create table if not exists polltype (
	id blob primary key, -- uuid v7
	name text unique not null
);

create table if not exists subject (
	id blob primary key, -- uuid v7
	name text unique not null
);

create table if not exists subject_polltype (
	subject_id blob,
	polltype_id blob,
	primary key (subject_id, polltype_id)
);
