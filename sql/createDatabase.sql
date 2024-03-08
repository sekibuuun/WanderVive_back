CREATE TABLE band (
    bandId SERIAL NOT NULL,
    bandName TEXT NOT NULL,
    genre TEXT default 'その他',
    youtube TEXT default 'NULL',
    twitter TEXT default 'NULL',
    instagram TEXT default 'NULL',
    tunecore TEXT default 'NULL',
    homePage TEXT default 'NULL',
    imagePath TEXT default 'NULL',
    primary key (bandId)
);

CREATE TABLE livehouse (
    livehouseId SERIAL NOT NULL,
    livehouseName TEXT NOT NULL,
    longitude NUMERIC(9, 7) NOT NULL ,
    latitude NUMERIC(10, 7) NOT NULL,
    homePage TEXT default 'NULL',
    mapLink TEXT default 'NULL',
    primary key(livehouseId)
);


CREATE TABLE event (
    eventId SERIAL NOT NULL,
    eventName TEXT NOT NULL,
    livehouseId integer,
    eventDate char(10) NOT NULL,
    openTime char(5) NOT NULL,
    startTime char(5) NOT NULL, 
    fee integer NOT NULL,
    primary key (eventId),
    foreign key (livehouseId) REFERENCES livehouse(livehouseId)
);

create table paticipant (
    Id serial not null,
    eventId integer not null,
    bandId integer not null,
    primary key(Id),
    foreign key(eventId) references event(eventId),
    foreign key(bandId) references event(eventId)
);