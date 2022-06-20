CREATE TABLE "blog" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "text" varchar NOT NULL,
  "image" bytea,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "question" (
  "id" bigserial PRIMARY KEY,
  "text" varchar NOT NULL,
  "answer_id" bigserial UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "answer" (
  "id" bigserial PRIMARY KEY,
  "answer_id" integer NOT NULL,
  "text" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "answer" ("id");

ALTER TABLE "answer" ADD FOREIGN KEY ("answer_id") REFERENCES "question" ("answer_id");

