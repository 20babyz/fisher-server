CREATE TABLE users (
     id INT AUTO_INCREMENT PRIMARY KEY,
     name VARCHAR(255) UNIQUE NOT NULL,
     image_url TEXT,
     created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);