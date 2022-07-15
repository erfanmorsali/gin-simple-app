create table users
(
    id       serial
        constraint users_pk
            primary key,
    username varchar(50)  not null,
    password varchar(100) not null,
    email    varchar(100) not null
);

create unique index users_username_uindex
    on users (username);

create unique index users_email_uindex
    on users (email);


create table posts
(
    id          serial
        constraint posts_pk
            primary key,
    title       varchar(50),
    description text,
    user_id     integer
        constraint posts___user_id_fk
            references users
);

create table categories
(
    id   serial
        constraint categories_pk
            primary key,
    name varchar(50)
);

create unique index categories_name_uindex
    on categories (name);


CREATE TABLE posts_categories
(
    post_id     int REFERENCES posts (id) ON UPDATE CASCADE ON DELETE CASCADE,
    category_id int REFERENCES categories (id) ON UPDATE CASCADE,
    CONSTRAINT posts_categories_pkey PRIMARY KEY (post_id, category_id)
);
