CREATE TABLE IF NOT EXISTS public.rooms
(
    id              SERIAL PRIMARY KEY,
    house_id INTEGER     REFERENCES  public.houses(id),
    name TEXT NOT NULL,
    description TEXT,
    created_date    timestamp NOT NULL,
    updated_date    timestamp NOT NULL,
    deleted_date    timestamp
);