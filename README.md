

sample structure

- api 
  - builder/ orchestrator
- routes
  - where sub routes go
    - these have controllers (db passed into)
    - services
    - models
- config
  - anything to do with getting or reading config
- db
  - setting up the db
  - migrations

design questions?
- should models be top level? will need to see if
we go with orm or sqlx