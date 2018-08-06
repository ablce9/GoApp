# zenport.io - Go assignment

This is the repository with the go assignment of **Zenport.io**.

Read the content of this **README.md** carefully as it will help you to set up the project
for this assignment and to submit your solution.

In the assignment, you will have to contribute a micro API with Golang.

## Getting started

### Requirements

The project needs few elements:

 * Go (obviously :D)
 * Go dependency management [dep](https://golang.github.io/dep/)
 * a **SQL database** (ex: PostgresSQL)
 * a **git** client to clone the sources

### Project setup

This repository uses git, you should clone it on your machine through a git clone in `$GOPATH/src/gitlab.com/zenport.io/go-assignment`.

Note: the project as no dependency for now so no need to run `dep ensure`.

After that your ready to go!

## It's up to you now!

### Instructions

Now read the [INSTRUCTIONS.md](./INSTRUCTIONS.md) file located at the root of the project. It contains all the instructions and hints for the project.

You can of course do several commits during the test to save your work. We'll look at the final result only.

### Questions

After that, read the [QUESTIONS.md](./QUESTIONS.md) file located at the root of the project. It contains some questions about the project or more.

### Submissions

You can submit notes in [NOTES.md](./NOTES.md)

To submit your solution, you can simply push your work into a repository of your choice that we can access.

## How you will be evaluated

### Tests

A bunch of tests have been written under `adapters/http` and `domain` (also in `providers/database`).

Those tests will be executed to see if your solution works.

Please justify if you did any modifications in tests.

### Code review

Your code will be also reviewed.

### Procedure how we test

 - We will clone your sources.
 - We will run `dep ensure`.
 - If needed, run `docker-compose up -d`.
 - We will run test like mention in [INSTRUCTIONS.md](./INSTRUCTIONS.md) or something else if specified.
 - We will run docker to build an image `docker build . -t go-assignment`

### About the results

The results are important but there is no unique solution, there are various way to pass the tests successfully. 
We prefer a not so perfect solution than no solution at all even if you have to change a the structure of the project.

**Don't forget that you will have to explain and justify your choices during an interview.
This exercise is mainly a way to provide discussion material for it.**
