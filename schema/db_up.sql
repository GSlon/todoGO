CREATE TABLE users
(
    id SERIAL NOT NULL UNIQUE,
    name VARCHAR(30) NOT NULL,
    surname VARCHAR(30) NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE todoitems
(
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(128) NOT NULL,
    description VARCHAR(255),
    user_id int REFERENCES users(id) ON DELETE CASCADE NOT NULL
);
