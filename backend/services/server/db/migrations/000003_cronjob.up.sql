CREATE TABLE cronjobs (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  name text,
  started_at timestamp with time zone,
  ended_at timestamp with time zone,
  new_articles_count integer,
  crawler_id SERIAL
);

ALTER TABLE
  cronjobs
ADD
  CONSTRAINT fk_cronjobs_crawler FOREIGN KEY (crawler_id) REFERENCES crawlers(id);