CREATE TABLE crawlers (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  source_link text,
  feed_link text UNIQUE,
  crawl_type text,
  article_div text,
  article_title text,
  article_description text,
  article_link text,
  article_published text,
  article_authors text,
  schedule text
);

CREATE TABLE articles_sources (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  link text,
  feed_link text UNIQUE,
  title text,
  description text,
  image text
);

CREATE TABLE articles (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  title text,
  description text,
  link text,
  published timestamp with time zone,
  authors text,
  articles_sources_id SERIAL
);

ALTER TABLE
  crawlers
ADD
  CONSTRAINT fk_crawlers_articlesource_feed_link FOREIGN KEY (feed_link) REFERENCES articles_sources(feed_link);

ALTER TABLE
  articles
ADD
  CONSTRAINT fk_articles_articlesource_id FOREIGN KEY (articles_sources_id) REFERENCES articles_sources(id);

CREATE UNIQUE INDEX idx_articles_id ON public.articles USING btree (id);

CREATE UNIQUE INDEX idx_crawlers_id ON public.crawlers USING btree (id);