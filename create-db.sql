CREATE TABLE IF NOT EXISTS `Category` (
    `name`      VARCHAR (120) NOT NULL,
    `timestamp` TIMESTAMP     NOT NULL,
    PRIMARY KEY (name)
);

CREATE TABLE IF NOT EXISTS `User` (
    `username`      VARCHAR (32)    NOT NULL,
    `email`         VARCHAR (120)   NOT NULL,
    `passwordhash`  BINARY (64)     NOT NULL,
    `reputation`    INT             DEFAULT '0',
    `role`          INT UNSIGNED    DEFAULT '0',
    `salt`          BINARY (32)     NOT NULL,
    PRIMARY KEY (username)
);

CREATE TABLE IF NOT EXISTS `Thread` (
    `id`            INT UNSIGNED    NOT NULL    AUTO_INCREMENT,
    `name`          VARCHAR (120)   NOT NULL,
    `username`      VARCHAR (120)   NOT NULL,
    `categoryname`  VARCHAR (120),

    PRIMARY KEY (id),
    FOREIGN KEY (username) REFERENCES User(username)
        ON UPDATE CASCADE
        ON DELETE CASCADE
    FOREIGN KEY (categoryname) REFERENCES Category(name)
        ON UPDATE CASCADE
        ON DELETE CASCADE

);

CREATE TABLE IF NOT EXISTS `Message` (
    `id`            INT UNSIGNED    NOT NULL    AUTO_INCREMENT,
    `message`       TEXT            NOT NULL,
    `timestamp`     TIMESTAMP       NOT NULL,
    `username`      VARCHAR (120)   NOT NULL,
    `threadId`      INT UNSIGNED    NOT NULL,
    `parentmessage` INT UNSIGNED,
    PRIMARY KEY (id),
    FOREIGN KEY (username) REFERENCES User(username)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    FOREIGN KEY (threadId) REFERENCES Thread(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    FOREIGN KEY (parentmessage) REFERENCES Message(id)
        ON UPDATE RESTRICT
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `ThreadMessages` (
    `threadId`      INT UNSIGNED    NOT NULL,
    `messageId`     INT UNSIGNED    NOT NULL,
    PRIMARY KEY (threadId, messageId),
    FOREIGN KEY (threadId) REFERENCES Thread(id)
        ON UPDATE RESTRICT
        ON DELETE CASCADE,
    FOREIGN KEY (messageId) REFERENCES Message(id)
        ON UPDATE RESTRICT
        ON DELETE CASCADE
);
