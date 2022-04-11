create database cabservice;

create table cabservice.booking_list(
booking_id int not null auto_increment primary key,
from_latitude DECIMAL(10,8) not NULL,
from_longitude DECIMAL(11,8) not null,
to_latitude DECIMAL(10,8) not NULL,
to_longitude DECIMAL(11,8) not null,
u_email varchar(40) not null,
cab_id varchar(20) not null,
status varchar(20) not null, ---completed/ongoing
date_of_booking varchar(30),
price float,
distance float
)

create table cabservice.live_location_with_status(
    latitude DECIMAL(10,8) not NULL,
    longitude DECIMAL(11,8) not null,
    cab_id varchar(40) not null,
    status varchar(20) not null
)