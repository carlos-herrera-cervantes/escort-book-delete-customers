CREATE TABLE "public"."profile" (
    "id" varchar(100) NOT NULL,
    "customer_id" varchar(100) NOT NULL UNIQUE,
    "first_name" varchar(100),
    "last_name" varchar(100),
    "email" varchar(100) NOT NULL,
    "phone_number" varchar(20),
    "gender" varchar(20),
    "birthdate" date,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);
