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

create table if not exists pollster (
	id blob primary key, -- uuid v7
	name text unique not null
);

create table if not exists poll (
	id blob primary key, -- uuid v7
	votehub_id text unique not null,
	poll_type text not null,
	sample_size integer,
	population text,
	url text not null,
	created_at text not null,
	start_date text not null,
	end_date text not null,
	pollster_id blob not null,
	seat_name text,
	internal integer not null,
	partisan text,
	subject_id blob not null
);

create table if not exists poll_answer (
	id blob primary key, -- uuid v7
	poll_id blob not null,
	choice text,
	pct integer,
	unique (poll_id, choice)
);

create table if not exists poll_sponsor (
	id blob primary key, -- uuid v7
	poll_id blob not null,
	sponsor text,
	unique (poll_id, sponsor)
);


