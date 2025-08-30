# FAQ

<details>
<summary>What's the difference between Next/Previous and Up/Down?</summary>

```
1_first_migration.up.sql           next ->  2_second_migration.up.sql      ...
1_first_migration.down.sql  <- previous     2_second_migration.down.sql    ...
```
</details>

## Why two separate files (up and down) for a migration?

It makes all of our lives easier. No new markup/syntax to learn for users and existing database utility tools continue to work as expected.

## How many migrations can migrate handle?

Way more than you can write.

## Are the table tests in migrate_test.go bloated?

Yes and no. There are duplicate test cases for sure but they don't hurt here. In fact the tests are very visual now and might help new users understand expected behaviors quickly. Migrate from version x to y and y is the last migration? Just check out the test for that particular case and know what's going on instantly.

## What is Docker being used for?

Only for testing. See [testing/docker.go](testing/docker.go)

## Why not just use docker-compose?

It doesn't give us enough runtime control for testing. We want to be able to bring up containers fast and whenever we want, not just once at the beginning of all tests.

## What does "dirty" database mean?

Before a migration runs, each database sets a dirty flag. Execution stops if a migration fails and the dirty state persists, which prevents attempts to run more migrations on top of a failed migration. You need to manually fix the error and then "force" the expected version.

## Do I need to create a table for tracking migration version used?

No, it is done automatically.

## I have got an error `Dirty database version 1. Fix and force version`. What should I do?
Keep calm and refer to [the getting started docs](GETTING_STARTED.md#forcing-your-database-version).
