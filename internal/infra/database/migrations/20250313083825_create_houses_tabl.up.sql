CREATE TABLE IF NOT EXISTS public.houses
(
    id              SERIAL PRIMARY KEY,
    user_id INTEGER     REFERENCES  public.users(id),
    name TEXT NOT NULL,
    description TEXT,
    city TEXT NOT NULL,
    address TEXT NOT NULL,
    Lat DOUBLE PRECISION NOT NULL,
    Lon DOUBLE PRECISION NOT NULL,
    created_date    timestamp NOT NULL,
    updated_date    timestamp NOT NULL,
    deleted_date    timestamp
);