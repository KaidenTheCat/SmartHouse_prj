CREATE TABLE IF NOT EXISTS public.measurements (
    id SERIAL PRIMARY KEY,
    device_uuid UUID NOT NULL,
    room_id INTEGER NOT NULL,
    value TEXT NOT NULL,
    created_date TIMESTAMPTZ NOT NULL
);