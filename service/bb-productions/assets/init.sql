CREATE DATABASE IF NOT EXISTS bb;
USE bb;

CREATE TABLE IF NOT EXISTS bb_rating (
    dish_name_id INT AUTO_INCREMENT PRIMARY KEY,
    dish_name VARCHAR(255),
    bb_rating INT
);

INSERT INTO bb_rating (dish_name, bb_rating) VALUES
('karubikuppa', 5),
('curry', 4),
('spaghetti', 3),
('meuniere', 4),
('sandwich', 3),
('salad', 1),
('smoothie', 2),
('yakitori', 4),
('yakiniku', 5);
