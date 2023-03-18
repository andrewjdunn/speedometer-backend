 /*DROP TABLE IF EXISTS speed_record;*/
CREATE TABLE speed_record (
  id INT AUTO_INCREMENT NOT NULL,
  timestamp DATETIME NOT NULL,
  latency INT NOT NULL,
  uploadspeed float(25) NOT NULL,
  downloadspeed float(25) NOT NULL,
  distance float(25) NOT NULL,
  pingok bit(1) NOT NULL,
  PRIMARY KEY (`id`)
);

