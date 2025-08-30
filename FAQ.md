# FAQ

<details>
<summary>What's the difference between Next/Previous and Up/Down?</summary>

```
1-first-migration.up.sql           next ->  2-second-migration.up.sql
1-first-migration.down.sql  <- previous     2-second-migration.down.sql
```

</details>

<details>
<summary>How many migrations can migrate handle?</summary>
Way more than you can write.
</details>

<details>
<summary>What does "dirty" database mean?</summary>
Before a migration runs, each database sets a dirty flag. Execution stops if a migration fails and the dirty state persists, which prevents attempts to run more migrations on top of a failed migration. You need to manually fix the error and then "force" the expected version.
</details>


<details>
<summary>Do I need to create a table for tracking migration version used?</summary>
No, it is done automatically.
</details>


<details>
<summary>I'm getting a "Dirty database version" error. What should I do?</summary>
Keep calm and refer to [the getting started docs](GETTING_STARTED.md#forcing-your-database-version).
</details>
