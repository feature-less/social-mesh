CREATE TABLE IF NOT EXISTS public.threads(
  "id" UUID PRIMARY KEY UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  "poster_id" UUID PRIMARY KEY UNIQUE NOT NULL,
  "created" _at TIMESTAMPTZ DEFAULT current_timestamp,
  "updated_at" TIMESTAMPTZ DEFAULT current_timestamp,
  "deleted_at" TIMESTAMPTZ DEFAULT NULL,
  "title" varchar(120) NULL,
  "body" VARCHAR(1024) NOT NULL,
  "images" VARCHAR [] NULL,
  "links" VARCHAR [] NULL,
)