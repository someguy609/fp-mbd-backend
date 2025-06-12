-- CREATE DATABASE golang_template;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Tabel Pengguna (Users)
CREATE TABLE users (
    user_id VARCHAR(15) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    role VARCHAR(10) NOT NULL, 
    contact_info VARCHAR(100),
    password VARCHAR(255) NOT NULL 
);

-- Tabel Proyek (Projects)
CREATE TABLE projects (
    project_id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    status VARCHAR(20) NOT NULL, 
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    categories VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Anggota Proyek (Project Members)
CREATE TABLE project_members (
    project_member_id SERIAL PRIMARY KEY,
    role_project VARCHAR(10) NOT NULL, 
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    users_user_id VARCHAR(15) NOT NULL,
    projects_project_id INT NOT NULL,
    FOREIGN KEY (users_user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (projects_project_id) REFERENCES projects(project_id) ON DELETE CASCADE
);

-- Tabel Milestone Proyek
CREATE TABLE milestones (
    milestone_id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    due_date TIMESTAMP,
    status VARCHAR(100) NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    projects_project_id INT NOT NULL,
    FOREIGN KEY (projects_project_id) REFERENCES projects(project_id) ON DELETE CASCADE
);

-- Tabel Dokumen
CREATE TABLE documents (
    document_id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    file_url VARCHAR(255) NOT NULL,
    document_type VARCHAR(100), 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    projects_project_id INT NOT NULL,
    users_user_id VARCHAR(15) NOT NULL,
    FOREIGN KEY (projects_project_id) REFERENCES projects(project_id) ON DELETE CASCADE,
    FOREIGN KEY (users_user_id) REFERENCES users(user_id) ON DELETE SET NULL
);