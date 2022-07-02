create table airlines (
	id int not null AUTO_INCREMENT,
  name varchar(255) not null,
  primary key (id)
);

create table route (
	id int not null AUTO_INCREMENT,
  airline_id int not null,
  flight_code varchar(255) not null,
  origin varchar(255) not null,
  destination varchar(255) not null,
  primary key (id),
  foreign key (airline_id) references airlines(id)
);

create table schedule (
	id int not null AUTO_INCREMENT,
  route_id int not null,
  `date` date,
  time_of_departure time,
  duration time,
  status varchar(255),
  primary key (id),
  foreign key (route_id) references route(id)
);