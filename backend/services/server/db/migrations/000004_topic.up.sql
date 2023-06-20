CREATE TABLE categories (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  name text UNIQUE,
  illustration text
);

CREATE TABLE topics (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  name text UNIQUE,
  category_id SERIAL
);

ALTER TABLE
  topics
ADD
  CONSTRAINT fk_topics_catogory_id FOREIGN KEY (category_id) REFERENCES categories(id);

ALTER TABLE articles_sources 
ADD COLUMN topic_id SERIAL;

ALTER TABLE
  articles_sources
ADD
  CONSTRAINT fk_articles_sources_topic_id FOREIGN KEY (topic_id) REFERENCES topics(id);

CREATE UNIQUE INDEX idx_articlessource_topic_id ON public.articles_sources USING btree (topic_id);

INSERT INTO
  categories(created_at, name)
values
  (current_timestamp, 'Others');

INSERT INTO
  topics(created_at, name, category_id)
values
  (current_timestamp, 'Others', 1);