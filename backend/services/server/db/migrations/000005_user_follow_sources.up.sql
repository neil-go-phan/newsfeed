CREATE TABLE follows (
  username text,
  articles_source_id SERIAL,
  PRIMARY KEY (username, articles_source_id)
);

CREATE TABLE reads (
  username text,
  article_id SERIAL,
  articles_source_id SERIAL,
  PRIMARY KEY (username, article_id)
);

ALTER TABLE
  follows
ADD
  CONSTRAINT fk_follows_username FOREIGN KEY (username) REFERENCES users(username);

ALTER TABLE
  follows
ADD
  CONSTRAINT fk_follows_articles_sources_id FOREIGN KEY (articles_source_id) REFERENCES articles_sources(id);

ALTER TABLE
  reads
ADD
  CONSTRAINT fk_reads_article_id FOREIGN KEY (article_id) REFERENCES articles(id);

ALTER TABLE
  reads
ADD
  CONSTRAINT fk_reads_username FOREIGN KEY (username) REFERENCES users(username);

ALTER TABLE
  reads
ADD
  CONSTRAINT fk_reads_articles_source_id FOREIGN KEY (articles_source_id) REFERENCES articles_sources(id);

CREATE UNIQUE INDEX idx_follows ON public.follows USING btree (articles_source_id, username);

CREATE INDEX idx_follows_user ON public.follows USING btree (username);

CREATE UNIQUE INDEX idx_read ON public.reads USING btree (username, article_id);

CREATE INDEX idx_read_articles_source_id ON public.reads USING btree (articles_source_id);