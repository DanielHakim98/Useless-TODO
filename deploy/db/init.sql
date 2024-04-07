-- Create schema
-- CREATE SCHEMA IF NOT EXISTS useless_todo;

-- Create table
CREATE TABLE IF NOT EXISTS todo_list(
    id SERIAL PRIMARY KEY,
    title VARCHAR(50),
    content TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Insert sample data
INSERT INTO todo_list (title, content)
VALUES
    ('Grocery Shopping', 'Buy milk'),
    ('Work Tasks', 'Schedule client calls'),
    ('Home Projects', 'Fix leaking faucet in the bathroom'),
    ('Fitness Routine', 'Jogging for 30 minutes'),
    ('Study Plan', 'Read chapter 5 of biology book');

