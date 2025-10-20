-- migrate:up
CREATE TABLE IF NOT EXISTS notes (
  id             BIGSERIAL PRIMARY KEY,
  title          VARCHAR(50),
  description    TEXT,
  image          VARCHAR(255),
  reminder_time  TIMESTAMPTZ,
  created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER trg_notes_updated_at
BEFORE UPDATE ON notes
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- migrate:down
DROP TRIGGER IF EXISTS trg_notes_updated_at ON notes;
DROP TABLE IF EXISTS notes;
