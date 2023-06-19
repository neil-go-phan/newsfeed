CREATE TABLE crawlers (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  source_link text UNIQUE,
  feed_link text UNIQUE,
  crawl_type text,
  article_div text,
  article_title text,
  article_description text,
  article_link text,
  article_authors text,
  schedule text DEFAULT '@daily' NOT NULL, 
  articles_source_id SERIAL
);

CREATE TABLE articles_sources (
  id SERIAL PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  link text UNIQUE,
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
  articles_source_id SERIAL
);

ALTER TABLE
  crawlers
ADD
  CONSTRAINT fk_crawlers_articlesource_source_link FOREIGN KEY (source_link) REFERENCES articles_sources(link);

ALTER TABLE
  crawlers
ADD
  CONSTRAINT fk_crawlers_articlesource_id FOREIGN KEY (articles_source_id) REFERENCES articles_sources(id);

ALTER TABLE
  articles
ADD
  CONSTRAINT fk_articles_articlesource_id FOREIGN KEY (articles_source_id) REFERENCES articles_sources(id);


CREATE UNIQUE INDEX idx_crawlers_id ON public.crawlers USING btree (id);

CREATE UNIQUE INDEX idx_articlessource_link ON public.articles_sources USING btree (link);
