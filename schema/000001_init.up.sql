CREATE TABLE users
(
    id serial not null unique,
    username varchar(255) not null,
    password_hash varchar(255) not null
);
CREATE TABLE notes
(
    id serial not null unique,
    description varchar(255) not null,
    user_id int references users(id) on delete cascade not null
);