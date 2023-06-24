Drop Table if exists data; 

CREATE TABLE data (
	id INTEGER PRIMARY KEY,
	submit_time VARCHAR(100),
	name VARCHAR(100),
	email VARCHAR(100)
	depature_city VARCHAR(100),
	depature_date VARCHAR(100),
    depature_time VARCHAR(100), 
    ariving_date VARCHAR(100), 
    ariving_time VARCHAR(100), 
    flight_num VARCHAR(100)
);

