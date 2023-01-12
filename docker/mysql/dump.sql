CREATE TABLE tables
(
    id        INT NOT NULL auto_increment,
    capacity  INT,
    available INT,
    PRIMARY KEY (id)
);

CREATE TABLE guests
(
    name         VARCHAR(255) UNICODE,
    table_id     INT,
    accompanying INT,
    time_Arrived DATETIME,
    PRIMARY KEY (name),
    FOREIGN KEY (table_id) REFERENCES tables (id)
);