CREATE SCHEMA IF NOT EXISTS ARTICLES;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE OR REPLACE FUNCTION ARTICLES.Trigger_SetTimestamp()  
RETURNS TRIGGER AS $$  
BEGIN  
  NEW.last_updated = NOW();
  RETURN NEW;
END;  
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION ARTICLES.ARRAY_DISTINCT(anyarray)
RETURNS anyarray AS $f$
  SELECT array_agg(DISTINCT x) FROM unnest($1) t(x);
$f$ LANGUAGE SQL IMMUTABLE;


CREATE OR REPLACE FUNCTION ARTICLES.ARRAY_DISTINCT_MINUS(anyarray, anyarray) RETURNS anyarray AS $f$
  SELECT array_agg(DISTINCT x) FROM 
	unnest(ARRAY(SELECT unnest($1) 
               EXCEPT 
               SELECT unnest($2))) t(x);
$f$ LANGUAGE SQL IMMUTABLE;

CREATE TABLE IF NOT EXISTS ARTICLES.ARTICLE(
    id serial NOT NULL,
    title VARCHAR (255) UNIQUE NOT NULL,
    publish_date DATE NOT NULL,
    body TEXT NOT NULL,
    tags TEXT[],
    created_on TIMESTAMP NOT NULL DEFAULT NOW(),
    last_updated TIMESTAMP,
    PRIMARY KEY (id)
);


CREATE TRIGGER ArticleSetTimestamp  
BEFORE UPDATE ON ARTICLES.ARTICLE  
FOR EACH ROW  
EXECUTE PROCEDURE ARTICLES.Trigger_SetTimestamp();

CREATE INDEX articles_tag_dex ON ARTICLES.ARTICLE USING gin(tags);