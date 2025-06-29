CREATE TABLE IF NOT EXISTS categories(
    id int PRIMARY KEY AUTO_INCREMENT,
    name varchar(255) NOT NULL
);


CREATE TABLE IF NOT EXISTS courses(
    id int PRIMARY KEY AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    category_id int NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);