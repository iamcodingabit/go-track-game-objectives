CREATE TABLE IF NOT EXISTS game(
    id INTEGER PRIMARY KEY,
    title TEXT UNIQUE NOT NULL,
    genre TEXT
);

CREATE TABLE IF NOT EXISTS platform(
    id INTEGER PRIMARY KEY,
    name TEXT UNIQUE NOT NULL 
);

CREATE TABLE IF NOT EXISTS playthrough(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    status TEXT NOT NULL DEFAULT 'backlog' CHECK(status IN ('backlog', 'in_progress', 'completed')),
    game_id INTEGER NOT NULL REFERENCES game(id) ON DELETE CASCADE,
    platform_id INTEGER NOT NULL REFERENCES platform(id),
    UNIQUE(game_id, platform_id, name)
);

CREATE INDEX IF NOT EXISTS idx_playthrough_platform ON playthrough(platform_id);

CREATE TABLE IF NOT EXISTS objective(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT,
    description TEXT,
    priority INTEGER NOT NULL DEFAULT 1 CHECK (priority BETWEEN 1 AND 5),
    is_completed INTEGER NOT NULL DEFAULT 0 CHECK (is_completed IN (0, 1)),
    parent_id INTEGER references objective(id) ON DELETE CASCADE,
    playthrough_id INTEGER NOT NULL REFERENCES playthrough(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_objective_parent ON objective(parent_id);

CREATE UNIQUE INDEX IF NOT EXISTS idx_objective_unique
ON objective(playthrough_id, COALESCE(parent_id, 0), name);