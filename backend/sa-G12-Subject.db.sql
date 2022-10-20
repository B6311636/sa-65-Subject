BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "officers" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"email"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "faculties" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "teachers" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"level"	text,
	"name"	text,
	"email"	text,
	"faculty_id"	integer,
	PRIMARY KEY("id"),
	CONSTRAINT "fk_faculties_teacher" FOREIGN KEY("faculty_id") REFERENCES "faculties"("id")
);
CREATE TABLE IF NOT EXISTS "times" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"period"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "subjects" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"code"	text,
	"name"	text,
	"section"	integer,
	"day"	text,
	"take"	integer,
	"teacher_id"	integer,
	"faculty_id"	integer,
	"officer_id"	integer,
	"time_id"	integer,
	PRIMARY KEY("id"),
	CONSTRAINT "fk_teachers_subject" FOREIGN KEY("teacher_id") REFERENCES "teachers"("id"),
	CONSTRAINT "fk_faculties_subject" FOREIGN KEY("faculty_id") REFERENCES "faculties"("id"),
	CONSTRAINT "fk_officers_subject" FOREIGN KEY("officer_id") REFERENCES "officers"("id"),
	CONSTRAINT "fk_times_subject" FOREIGN KEY("time_id") REFERENCES "times"("id")
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_officers_email" ON "officers" (
	"email"
);
CREATE INDEX IF NOT EXISTS "idx_officers_deleted_at" ON "officers" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_faculties_deleted_at" ON "faculties" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_teachers_deleted_at" ON "teachers" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_times_deleted_at" ON "times" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_subjects_deleted_at" ON "subjects" (
	"deleted_at"
);
COMMIT;
