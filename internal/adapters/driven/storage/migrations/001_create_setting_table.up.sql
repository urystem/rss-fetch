CREATE TABLE IF NOT EXISTS setting (
    mine BOOLEAN PRIMARY KEY DEFAULT TRUE,       -- гарантирует 1 запись
    is_running BOOLEAN NOT NULL DEFAULT FALSE, -- статус fetch
    interval INTERVAL NOT NULL DEFAULT '3 minutes',
    workers INT NOT NULL DEFAULT 3 CHECK (workers > 0)
);
INSERT INTO setting DEFAULT VALUES;

CREATE OR REPLACE FUNCTION notify_setting_change() RETURNS trigger AS $$
BEGIN
    PERFORM pg_notify('setting_changed', TG_OP);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER setting_table_change_trigger
AFTER UPDATE OR DELETE ON setting
FOR EACH ROW EXECUTE FUNCTION notify_setting_change();
