-- +goose Up
-- SQL in this section is executed when the migration is applied.

create table booking_list(
booking_id int not null auto_increment primary key,
from_latitude DECIMAL(10,8) not NULL,
from_longitude DECIMAL(11,8) not null,
to_latitude DECIMAL(10,8) not NULL,
to_longitude DECIMAL(11,8) not null,
u_email varchar(40) not null,
cab_id varchar(20) not null,
status varchar(20) not null, 
date_of_booking varchar(30),
price float,
distance float
);

create table live_location_with_status(
    latitude DECIMAL(10,8) not NULL,
    longitude DECIMAL(11,8) not null,
    cab_id varchar(40) not null,
    status varchar(20) not null
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

drop table booking_list;
drop table live_location_with_status;

