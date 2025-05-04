PRAGMA foreign_keys = ON;



CREATE TABLE users IF NOT EXISTS (
    userID INTEGER PRIMARY KEY AUTO_INCREMENT, 
    nickname VARCHAR(30) NOT NULL UNIQUE, 
    age INTEGER NOT NULL, 
    gender TEXT CHECK (gender IN ("Male","Female")) NOT NULL DEFAULT "Male",
    firstName VARCHAR(30) NOT NULL,
    lastName VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(50) NOT NULL,

)WITHOUT ROWID;


CREATE TABLE sessions IF NOT EXISTS (
    sessionID INTEGER PRIMARY KEY AUTO_INCREMENT,
    userID INTEGER NOT NULL UNIQUE,
    sessionToken VARCHAR(200) NOT NULL,
    expiresAt DATETIME,
    FOREIGN KEY (userID) REFERENCES users(userID) ON DELETE CASCADE,
); 


CREATE TABLE posts IF NOT EXISTS (
  postID INTEGER PRIMARY KEY AUTO_INCREMENT,
  userID INTEGER NOT NULL, 
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  total_comments INTEGER CHECK (total_comments >= 0) DEFAULT 0,  
  FOREIGN KEY (userID) REFERENCES users(userID) ON DELETE CASCADE 

);

CREATE TABLE comments IF NOT EXISTS  (
        commentID INTEGER PRIMARY KEY AUTOINCREMENT,
        postID INT,
        userID INT,
        createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        content TEXT NOT NULL,
        FOREIGN KEY (postID) REFERENCES posts (postID) ON DELETE CASCADE,
        FOREIGN KEY (userID) REFERENCES users (userID) ON DELETE CASCADE,
    );

CREATE TABLE categories IF NOT EXISTS (
    categorieID INTEGER PRIMARY KEY AUTO_INCREMENT,
    category TEXT NOT NULL,
);



INSERT INTO categories (category) VALUES (
  ("Discussions"),
  ("Questions"),
  ("Ideas"),
  ("Articles"),
  ("Events"), 
  ("Issues")

);