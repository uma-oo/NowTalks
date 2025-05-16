PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS users (
    userID INTEGER PRIMARY KEY AUTOINCREMENT,
    nickname VARCHAR(30) NOT NULL UNIQUE,
    age INTEGER NOT NULL,
    gender TEXT NOT NULL DEFAULT 'Male' CHECK (gender IN ('Male','Female')),
    firstName VARCHAR(30) NOT NULL,
    lastName VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(50) NOT NULL, 
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sessions (
    sessionID INTEGER PRIMARY KEY,
    userID  INTEGER NOT NULL UNIQUE,
    sessionToken VARCHAR(200) NOT NULL UNIQUE,
    expiresAt DATETIME,
    FOREIGN KEY (userID) REFERENCES users(userID) ON DELETE CASCADE
); 

CREATE TABLE IF NOT EXISTS posts (
    postID INTEGER PRIMARY KEY AUTOINCREMENT,
    userID INTEGER NOT NULL, 
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    total_comments INTEGER DEFAULT 0 CHECK (total_comments >= 0),  
    FOREIGN KEY (userID) REFERENCES users(userID) ON DELETE CASCADE 
);

CREATE TABLE IF NOT EXISTS comments (
    commentID INTEGER PRIMARY KEY,
    postID INTEGER NOT NULL,
    userID INTEGER NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    content TEXT NOT NULL ,
    FOREIGN KEY (postID) REFERENCES posts(postID) ON DELETE CASCADE,
    FOREIGN KEY (userID) REFERENCES users(userID) ON DELETE CASCADE
);


DROP TABLE IF EXISTS categories ;

CREATE TABLE IF NOT EXISTS categories (
    categoryID INTEGER PRIMARY KEY,
    category TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS postCategories (
    categoryID INTEGER NOT NULL,
    postID INTEGER NOT NULL,
    PRIMARY KEY (categoryID, postID)
    FOREIGN KEY (postID) REFERENCES posts(postID) ON DELETE CASCADE,
    FOREIGN KEY (categoryID) REFERENCES categories(categoryID) ON UPDATE CASCADE
);



INSERT INTO categories (category) VALUES
  ('Discussions'),
  ('Questions'),
  ('Ideas'),
  ('Articles'),
  ('Events'), 
  ('Issues');


/* */


