CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255),
    phone VARCHAR(20),
    email VARCHAR(255),
    subscription_lvl INT,
    current_gym_id INT
);