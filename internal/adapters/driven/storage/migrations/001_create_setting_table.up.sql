CREATE TABLE IF NOT EXISTS setting (
    mine BOOLEAN PRIMARY KEY DEFAULT TRUE,       -- гарантирует 1 запись //only for 'where'
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

-- триггер запрещает изменение mine
-- Универсальный триггер
CREATE OR REPLACE FUNCTION protect_mine()
RETURNS trigger AS $$
BEGIN
    -- Запрещаем изменение mine
    IF TG_OP = 'UPDATE' AND OLD.mine <> NEW.mine THEN
        RAISE EXCEPTION 'Cannot update column mine';
    END IF;

    -- Запрещаем удаление строки
    IF TG_OP = 'DELETE' THEN
        RAISE EXCEPTION 'Cannot delete setting row';
    END IF;
     
    IF TG_OP = 'INSERT' THEN
        RAISE EXCEPTION 'Cannot insert setting row';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Создаём триггер для всех операций
CREATE TRIGGER protect_mine_trigger
BEFORE INSERT OR UPDATE OR DELETE ON setting
FOR EACH ROW
EXECUTE FUNCTION protect_mine();