-- Create schema
CREATE SCHEMA IF NOT EXISTS useless_todo;

-- Create table
CREATE TABLE IF NOT EXISTS useless_todo.todo_list(
    id SERIAL PRIMARY KEY,
    title VARCHAR(50),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Insert sample data
INSERT INTO useless_todo.todo_list (title, content) VALUES
    ('Task 1', 'This is the content for task 1.'),
    ('Task 2', 'This is the content for task 2.'),
    ('Task 3', 'This is the content for task 3.');

