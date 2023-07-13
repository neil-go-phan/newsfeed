CREATE TABLE fcm_notifications (
  username text,
  firebase_token text,
  PRIMARY KEY (username, firebase_token)
);

ALTER TABLE
  fcm_notifications
ADD
  CONSTRAINT fk_fcm_notifications_username FOREIGN KEY(username) REFERENCES users(username) on delete cascade;