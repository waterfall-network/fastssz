
# Contributing to FastSSZ

Bear in mind that code generation is not straighforward, specially when you do not use something like custom langs and so on. Thus, some code generation is not easy to do and introduces much complexity.

Specially things like remote links and so on are not hard to do and we do it in an opinionated way that works so far. Consider that this was done on the scope of the Prysm project and it is preferable if you adapt to the way they do things.

For example, things like:
- alias of an alias.
- Not using pointers.
- Complex external dependencies.

## Codebase

This is how the codebase is separated:

- / : Helper methods shared among all the generate functions (i.e. encode an int to bytes or merkle hash an array).
- /sszgen: SSZ generator.
- /spectests: Representation of the official Eth2.0 structs.
- /fuzz: A framework to do fuzzing on the spectests (not production ready).

## Create an issue

First create an issue to propose the new change such that we can discuss how to approach the problem.

## Testing

We use the spectests repo as an e2e testing.
After each new change and before making a PR, make sure that you regenerate all the spectests again, this is twofold:
- We can test with the specs if everything work.
- We can visually see in the Github diff what parts have been updated in the branch.

Note that the CI (working on it) will fail if the spectests are not changed.

## Tips to update it

Keep everything simple, bound check should be done in the generated structs and not in the shared functions.
Fastssz is meant to be a high performant library so everything related to variable declaration should be avoided in the generate structs. If you find yourself using ":=" out of intergers and loop ranging you need to check again.
Note that this library is meant to be used on Prysm

If you are doing a lot of stuff rather than bound check and looping, you should create a function in the root.
