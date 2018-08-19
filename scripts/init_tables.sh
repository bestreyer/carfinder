#!/usr/bin/env bash

psql ${POSTGRES_DBNAME} ${POSTGRES_USER} << EOF
CREATE TABLE IF NOT EXISTS driver_location
(
  driver_id serial PRIMARY KEY,
  location geography(POINT),
  latitude double precision,
  longitude double precision,
  updated_at timestamp without time zone
);

CREATE INDEX IF NOT EXISTS driver_location_position_gix ON public.driver_location USING GIST (location);
EOF