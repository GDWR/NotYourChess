CREATE VIEW random_match AS
  SELECT *
  FROM match
  ORDER BY random();