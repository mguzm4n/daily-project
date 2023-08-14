UPDATE notes SET user_id = NULL;
ALTER TABLE notes DROP CONSTRAINT IF EXISTS notes_user_id_fkey;
ALTER TABLE notes DROP COLUMN IF EXISTS user_id;

-- Drop the users table
DROP TABLE IF EXISTS users;