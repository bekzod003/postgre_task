CREATE DATABASE udevs;

CREATE TABLE task_list (
    id serial primary key,
    name varchar not null,
    status varchar not null,
    created_at date not null default now(),
    created_by varchar not null,
    due_date date
);

CREATE TABLE contact_list (
    id serial primary key,
    first_name varchar not null,
    last_name varchar not null,
    phone varchar not null,
    email varchar not null
);