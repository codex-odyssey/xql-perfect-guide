CREATE DATABASE IF NOT EXISTS bbb;
USE bbb;

CREATE TABLE IF NOT EXISTS bbb_evaluation (
    dish_name_id INT AUTO_INCREMENT PRIMARY KEY,
    dish_name VARCHAR(255),
    bbb_evaluation INT
);

INSERT INTO bbb_evaluation (dish_name, bbb_evaluation) VALUES
('karubikuppa', 4);