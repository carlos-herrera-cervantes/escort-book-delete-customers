CREATE TABLE "public"."nationality" (
    "id" varchar(100) NOT NULL,
    "name" varchar(100) NOT NULL,
    "active" bool NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."profile" (
    "id" varchar(100) NOT NULL,
    "escort_id" varchar(100) NOT NULL UNIQUE,
    "first_name" varchar(50),
    "last_name" varchar(50),
    "email" varchar(100) NOT NULL,
    "phone_number" varchar(20),
    "gender" varchar(15),
    "nationality_id" varchar(100),
    "birthdate" date,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    CONSTRAINT "profile_nationality_id_fkey" FOREIGN KEY ("nationality_id") REFERENCES "public"."nationality"("id"),
    PRIMARY KEY ("id")
);
