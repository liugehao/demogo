CREATE TABLE origin.scsj
(
  id       SERIAL8,
  category CHAR(3),
  sta      VARCHAR(20),
  sta_tm   TIMESTAMP,
  batt     FLOAT,
  ver      VARCHAR(10),
  sn       VARCHAR(20),
  sn_tm    TIMESTAMP,
  pressure FLOAT,
  v_up     FLOAT,
  v_down   FLOAT,
  err      VARCHAR(10),
  tm       TIMESTAMP
);