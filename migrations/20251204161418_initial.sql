-- Create "opinions" table
CREATE TABLE "public"."opinions" (
  "id" bigserial NOT NULL,
  "user_id" bigint NOT NULL,
  "subject_type" text NOT NULL,
  "subject_id" bigint NOT NULL,
  "opinion" text NOT NULL,
  "rating" integer NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "opinions_rating_check" CHECK ((rating >= 1) AND (rating <= 5))
);
-- Create index "idx_opinions_subject" to table: "opinions"
CREATE INDEX "idx_opinions_subject" ON "public"."opinions" ("subject_type", "subject_id");
-- Create index "idx_opinions_user" to table: "opinions"
CREATE INDEX "idx_opinions_user" ON "public"."opinions" ("user_id");
