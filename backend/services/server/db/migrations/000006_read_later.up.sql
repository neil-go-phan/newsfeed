CREATE TABLE read_laters (
  username text,
  article_id SERIAL,
  PRIMARY KEY (username, article_id)
);


ALTER TABLE
  read_laters
ADD
  CONSTRAINT fk_read_laters_article_id FOREIGN KEY (article_id) REFERENCES articles(id);

ALTER TABLE
  read_laters
ADD
  CONSTRAINT fk_read_laters_username FOREIGN KEY (username) REFERENCES users(username);

CREATE INDEX idx_read_laters_username ON public.read_laters USING btree (username);
