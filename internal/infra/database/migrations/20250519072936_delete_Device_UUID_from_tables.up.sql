ALTER TABLE public.measurements
DROP COLUMN IF EXISTS device_uuid;

ALTER TABLE public.events
DROP COLUMN IF EXISTS device_uuid;

ALTER TABLE public.measurements
ADD COLUMN IF NOT EXISTS device_id INTEGER;

ALTER TABLE public.events
ADD COLUMN IF NOT EXISTS device_id INTEGER;
