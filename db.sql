CREATE DATABASE db
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Russian_Russia.1251'
    LC_CTYPE = 'Russian_Russia.1251'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;
CREATE TABLE public.room
(
    room_id integer NOT NULL,
    price integer NOT NULL,
    date date NOT NULL
)

TABLESPACE pg_default;

ALTER TABLE public.room
    OWNER to postgres;
CREATE TABLE public.bookings
(
    booking_id integer NOT NULL DEFAULT nextval('bookings_booking_id_seq'::regclass),
    room_id integer NOT NULL,
    data_start date NOT NULL,
    data_end date NOT NULL,
    CONSTRAINT bookings_pkey PRIMARY KEY (booking_id)
)

TABLESPACE pg_default;

ALTER TABLE public.bookings
    OWNER to postgres;
