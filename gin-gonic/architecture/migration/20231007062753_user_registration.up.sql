CREATE TABLE user_registration(
    name varchar(50) NOT NULL Unique,
    email varchar(50) NOT NULL Unique,
    password varchar(50) NOT NULL,
    phone varchar(20) NOT NULL,
    age int NOT NULL,
    address varchar(50) NOT NULL
);