CREATE DATABASE IF NOT EXISTS chef;
USE chef;

CREATE TABLE IF NOT EXISTS cooking_time (
    dish_name_id INT AUTO_INCREMENT PRIMARY KEY,
    dish_name VARCHAR(255),
    cooking_time INT
);

INSERT INTO cooking_time (dish_name, cooking_time) VALUES
('karubikuppa', 30),
('curry', 50),
('spaghetti', 20),
('meuniere', 30),
('sandwich', 10),
('salad', 5),
('smoothie', 15),
('yakitori', 10);