CREATE KEYSPACE test_keyspace WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1};
CREATE TABLE test_keyspace.url_redirect (url_hash text PRIMARY KEY, original_url text, url_id int);
