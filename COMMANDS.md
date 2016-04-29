# Algorithm


## code-review request <branch>

Creates a merge request from `<branch>` to `develop`. Ignores in case an open
merge request already exists.


## code-review accept <branch>

Merge and close a merge request to `develop`.

- Open merge request to `develop`;


## qa request <branch>

Creates a merge request from `<branch>` to `qa`. Ignores in case an open merge
request already exists.


## qa start <branch>

Merge and close a merge request to `qa`.

- Open merge request to `qa`;
- The HEAD of `qa` must be the same as the HEAD of `master` for all repositories
  of `<branch>` (there is no other feature in test);


## qa approve <branch>

Merge `qa` with `master` for all repositories of `<branch>` and delete
`<branch>`.

- The HEAD of `qa` must be the HEAD of `<branch>`, otherwise there is something
  untested in `<branch>`;
- `<branch>` must still exist;


## qa reject <branch>

Revert `qa` to the last commit before merge with `<branch>`.

- The HEAD of `qa` must be a merge with `<branch>`;


# List of commands

```
gpf code-review request "my-feature"
gpf code-review accept "my-feature"
gpf qa request "my-feature"
gpf qa start "my-feature"
gpf qa approve "my-feature"
gpf qa refuse "my-feature"
```

