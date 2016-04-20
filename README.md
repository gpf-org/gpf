gpf - git pipeline flow
=======================

# process

## branches

- develop (integration environment)
- qa (test environment)
- master (production environment - main)

## assumptions

- if commit of `master` is different from `qa`, it means there is something
  being tested in `qa`;
- if there is an open merge request to `develop` it means it is available for
  code reviewing;
- if there is an open merge request to `qa` it means it is available for
  testing;
- if a branch different from the ones listed in `branches` exist it means
  there is a working in progress (it may be finished and waiting for test
  or code reviewing;

## triggers

- whenever a new commit is pushed the CI is going to build and push to docker
  registry. the docker image will be tagged as following:
  - develop -> develop
  - qa -> qa
  - master -> latest

- whenever a new `qa` tag is generated, it will be deployed in both `qa` and
  `staging` environments.

# cli

the cli is going to help the team to deal with the process described above.
the cli will support microservices development by grouping branches from
different projects that have the same name (it is important to define
different name in different projects in case they are not part of the same
user story).

although the tool is going to support microservices it is not limited to
it, it may be perfectly used by traditional software architectures.

the cli might be called: `gpf - git pipeline flow` or a variant of it.

## commands

```sh
gpf list [status]       # list feature branches and their statuses
gpf create <name>       # create a feature branch from master and checkout to it

gpf list code-review    # list features waiting and during code review
gpf code-review request # create merge request to develop
gpf code-review approve # approve and close merge request
gpf code-review deny    # deny and close merge request

gpf list test               # list features waiting and during test
gpf test <feature>          # merge feature with qa
gpf test <feature> approve  # merge qa with master
gpf test <feature> deny     # revert last merge
```
