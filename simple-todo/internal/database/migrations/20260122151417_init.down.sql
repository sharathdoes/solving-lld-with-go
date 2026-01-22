-- =========================
-- USERS
-- =========================
CREATE TABLE users (
    id TEXT PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- =========================
-- PROJECTS
-- =========================
CREATE TABLE projects (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT,

    owner_id TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_projects_owner
        FOREIGN KEY (owner_id)
        REFERENCES users(id)
        ON DELETE SET NULL
);

CREATE INDEX idx_projects_deleted_at ON projects(deleted_at);
CREATE INDEX idx_projects_owner_id ON projects(owner_id);

-- =========================
-- TASKS
-- =========================
CREATE TABLE tasks (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT,

    project_id TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_tasks_project
        FOREIGN KEY (project_id)
        REFERENCES projects(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_tasks_deleted_at ON tasks(deleted_at);
CREATE INDEX idx_tasks_project_id ON tasks(project_id);

-- =========================
-- PROJECT MEMBERS (many-to-many)
-- =========================
CREATE TABLE project_members (
    project_id TEXT NOT NULL,
    user_id TEXT NOT NULL,

    PRIMARY KEY (project_id, user_id),

    CONSTRAINT fk_project_members_project
        FOREIGN KEY (project_id)
        REFERENCES projects(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_project_members_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-- =========================
-- TASK ASSIGNEES (many-to-many)
-- =========================
CREATE TABLE task_assignees (
    task_id TEXT NOT NULL,
    user_id TEXT NOT NULL,

    PRIMARY KEY (task_id, user_id),

    CONSTRAINT fk_task_assignees_task
        FOREIGN KEY (task_id)
        REFERENCES tasks(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_task_assignees_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-- =========================
-- REFRESH TOKEN TTLS
-- =========================
CREATE TABLE refresh_token_ttls (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    token_hash TEXT NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    revoked BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT fk_refresh_tokens_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_refresh_token_ttls_user_id ON refresh_token_ttls(user_id);
