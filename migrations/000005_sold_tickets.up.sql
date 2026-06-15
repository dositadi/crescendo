BEGIN;

CREATE TYPE ticket AS ENUM ('General Admission','VIP Standing','Reserved Seated');

CREATE TABLE IF NOT EXISTS sold_tickets (
    id uuid NOT NULL,
    userId uuid NOT NULL,
    userContactFName text NOT NULL,
    userContactLName text NOT NULL,
    userContactEmail text NOT NULL, 
    artistId integer NOT NULL,
    concertDate text NOT NULL,
    ticketType ticket NOT NULL DEFAULT 'General Admission',
    qty integer NOT NULL DEFAULT 1,
    vat double precision NOT NULL,
    amt double precision NOT NULL, 
    location text NOT NULL, 
    bookingFee double precision NOT NULL,
    version integer NOT NULL DEFAULT 1,
    createdAt TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updateAt TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),

    CONSTRAINT fk_user FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX date_idx ON sold_tickets (concertDate);
CREATE INDEX location_idx ON sold_tickets (location);

COMMIT;