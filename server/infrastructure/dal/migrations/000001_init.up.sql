create table url
(
	hash varchar not null
		constraint url_pk
			primary key,
	initial_url text not null
);