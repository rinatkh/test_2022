create table pets
(
    name       text,
    age        integer,
    type       text,
    created_at timestamp,
    id         text not null
        constraint id
            primary key
);

alter table pets
    owner to postgres;

create table photo
(
    id         text not null
        primary key,
    url        text,
    created_at timestamp,
    pets_id    text
        constraint pets_id
            references pets
);

alter table photo
    owner to postgres;

