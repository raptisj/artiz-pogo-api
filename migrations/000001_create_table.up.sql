CREATE TABLE temp (data jsonb);

\COPY temp (data) FROM './artists_file.json';

name, birth_year, death_year, instrument, genre, bio
SELECT data->>'id', data->>'name', data->>'birth_year', data->>'death_year', data->>'instrument', data->>'genre', data->>'bio' 
FROM temp;

INSERT INTO artists SELECT (data->>'id')::integer, data->>'name', (data->>'birth_year')::integer, (data->>'death_year')::integer, data->>'instrument', data->>'genre', data->>'bio' FROM temp;

DROP TABLE temp;