CREATE TABLE IF NOT EXISTS public.events (
    id SERIAL PRIMARY KEY,
    device_uuid UUID NOT NULL,
    room_id INTEGER NOT NULL,
    action TEXT NOT NULL,
    created_date TIMESTAMPTZ NOT NULL
);