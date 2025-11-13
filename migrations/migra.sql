CREATE TABLE IF NOT EXISTS "sessions"
(
	session_id TEST PRIMARY KEY,
	user_id INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS "users"
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    "login" TEXT NOT NULL UNIQUE, 
	pas_hash TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS topics (
	name TEXT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS "message_main_table"
(
    name TEXT, 
    message TEXT, 
    date TEXT
);
