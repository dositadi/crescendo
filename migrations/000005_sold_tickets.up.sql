BEGIN;

CREATE TYPE ticket AS ENUM ('General Admission','VIP Standing','Reserved Seated');

CREATE TABLE IF NOT EXISTS sold_tickets (
    id uuid NOT NULL,
    userId uuid NOT NULL, 
    artistId integer NOT NULL,
    concertDate text NOT NULL,
    ticketType ticket NOT NULL DEFAULT 'General Admission',
    qty integer NOT NULL DEFAULT 1,
    vat double precision NOT NULL,
    amt double precision NOT NULL, 
    location text NOT NULL, 
    bookingFee double precision NOT NULL,
    version integer NOT NULL DEFAULT 1,

    CONSTRAINT fk_user FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX date_idx ON sold_tickets (concertDate);
CREATE INDEX location_idx ON sold_tickets (location);

COMMIT;