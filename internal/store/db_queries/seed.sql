INSERT OR IGNORE INTO game(title, genre) VALUES
("Dragon's Dogma 2: Dark Arisen", "Action RPG");
INSERT OR IGNORE INTO game(title, genre) VALUES
("Monster Hunter Wilds", "Action RPG");

INSERT OR IGNORE INTO platform(name) VALUES
("Steam");
INSERT OR IGNORE INTO platform(name) VALUES
("Nintendo Switch");

/*
id INTEGER PRIMARY KEY,
name TEXT NOT NULL,
description TEXT,
status TEXT NOT NULL DEFAULT 'backlog' CHECK(status IN ('backlog', 'in_progress', 'completed')),
game_id INTEGER NOT NULL REFERENCES game(id),
platform_id INTEGER NOT NULL REFERENCES platform(id)
*/

INSERT OR IGNORE INTO playthrough(name, description, status, game_id, platform_id)
SELECT 'Dark Arisen Expansion', '1st Playthrough', 'in_progress', game.id, platform.id
FROM game, platform
WHERE game.title = "Dragon's Dogma 2: Dark Arisen" AND
platform.name = "Steam";

/*
id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT,
    description TEXT,
    priority INTEGER NOT NULL DEFAULT 1 CHECK (priority BETWEEN 1 AND 5),
    is_completed INTEGER NOT NULL DEFAULT 0 CHECK (is_completed IN (0, 1)),
    parent_id INTEGER references objective(id) ON DELETE CASCADE,
    playthrough_id INTEGER NOT NULL REFERENCES playthrough(id) ON DELETE CASCADE
*/
INSERT OR IGNORE INTO objective(name, type, description, priority, is_completed, parent_id, playthrough_id)
SELECT 'Max out Warrior', 'Levelling', 'Hunt packs in Norgan', 1, 0, NULL, playthrough.id
FROM playthrough
JOIN game ON playthrough.game_id = game.id
WHERE game.title = 'Dragon''s Dogma 2: Dark Arisen'
AND playthrough.name = 'Dark Arisen Expansion';

INSERT OR IGNORE INTO objective (name, type, description, priority, is_completed, parent_id, playthrough_id)
SELECT 'Hunt southern pack', 'Fight', 'Hunt the pack south from the hub', 1, 0, parent.id, playthrough.id
FROM objective AS parent
JOIN playthrough ON parent.playthrough_id = playthrough.id
JOIN game ON playthrough.game_id = game.id
WHERE game.title = 'Dragon''s Dogma 2: Dark Arisen'
  AND playthrough.name = 'Dark Arisen Expansion'
  AND parent.name = 'Max out Warrior';