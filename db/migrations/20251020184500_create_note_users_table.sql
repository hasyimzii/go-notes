-- migrate:up
CREATE TABLE IF NOT EXISTS note_users (
  id          BIGSERIAL PRIMARY KEY,
  note_id     BIGINT NOT NULL REFERENCES notes(id) ON DELETE CASCADE,
  user_id     BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  CONSTRAINT uniq_note_users UNIQUE (note_id, user_id)
);

CREATE INDEX idx_note_users_user_id ON note_users(user_id);
CREATE INDEX idx_note_users_note_id ON note_users(note_id);

CREATE TRIGGER trg_note_users_updated_at
BEFORE UPDATE ON note_users
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- migrate:down
DROP TRIGGER IF EXISTS trg_note_users_updated_at ON note_users;
DROP TABLE IF EXISTS note_users;
