### How to store latitude and longitude
`driver_location` table has three columns, which to store location: `latitude`, `longitude`, `location`

`location` column is geometry type column. latitude and longitude columns are double precision type column

I use geometry type column, because we can assume that within 50 meters the area is flat. Also geometry types have better perfomance than geography types.
I store both location and pair (latitude, longitude), because ST_AsText function, which converts from geometry to latitude, latitude is slow function

```
CREATE TABLE driver_location
(
  driver_id serial PRIMARY KEY,
  location geometry(POINT, 4326),
  latitude double precision,
  longitude double precision,
  updated_at timestamp without time zone
);

CREATE INDEX driver_location_position_gix ON public.driver_location USING GIST (location);
```