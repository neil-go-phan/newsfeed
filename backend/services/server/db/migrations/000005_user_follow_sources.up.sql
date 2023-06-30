CREATE TABLE follows (
  created_at timestamp with time zone default current_timestamp,
  username text NOT NULL,
  articles_source_id SERIAL NOT NULL,
  unread integer default 0 NOT NULL,
  PRIMARY KEY (username, articles_source_id)
);

CREATE TABLE reads (
  created_at timestamp with time zone default current_timestamp,
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

CREATE INDEX idx_follows_created_at ON public.follows USING btree (created_at);

CREATE INDEX idx_read_username ON public.reads USING btree (username);

CREATE INDEX idx_read_username_articles_source_id ON public.reads USING btree (username, articles_source_id);

CREATE OR REPLACE FUNCTION decrease_unread()
  RETURNS TRIGGER AS $$
BEGIN
  IF NEW.username IS NOT NULL THEN
    UPDATE follows
    SET unread = unread - 1
    WHERE username = NEW.username
      AND articles_source_id = NEW.articles_source_id
      AND unread > 0;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_decrease_unread
AFTER INSERT ON reads
FOR EACH ROW
EXECUTE FUNCTION decrease_unread();

CREATE OR REPLACE FUNCTION increase_unread()
  RETURNS TRIGGER AS $$
BEGIN
  IF OLD.username IS NOT NULL THEN
    UPDATE follows
    SET unread = unread + 1
    WHERE username = OLD.username
      AND articles_source_id = OLD.articles_source_id;
  END IF;
  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_increase_unread
AFTER DELETE ON reads
FOR EACH ROW
EXECUTE FUNCTION increase_unread();