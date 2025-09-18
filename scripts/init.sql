-- Initialize database schema for goscaffold

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_created_at ON users(created_at);

-- Create a function to update the updated_at column
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create trigger to automatically update updated_at
CREATE TRIGGER update_users_updated_at 
    BEFORE UPDATE ON users 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Insert sample data for development
INSERT INTO users (id, email, first_name, last_name) VALUES
    ('550e8400-e29b-41d4-a716-446655440001', 'john.doe@example.com', 'John', 'Doe'),
    ('550e8400-e29b-41d4-a716-446655440002', 'jane.smith@example.com', 'Jane', 'Smith'),
    ('550e8400-e29b-41d4-a716-446655440003', 'bob.johnson@example.com', 'Bob', 'Johnson')
ON CONFLICT (email) DO NOTHING;
