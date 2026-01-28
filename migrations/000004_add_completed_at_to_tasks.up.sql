
ALTER TABLE tasks 
ADD COLUMN IF NOT EXISTS completed_at TIMESTAMP;

UPDATE tasks 
SET completed_at = updated_at 
WHERE completed = true AND completed_at IS NULL;