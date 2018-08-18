# Contributing
Thank you for your interest in contributing to out project. The following are guidelines on how to best contribute to development of the HackIllinois API. Any contibution, wheather it be a bug report, feature request, bug fix, feature addition, or additional documentation, is greatly appreciated.

1. [How to Contribute](#how-to-contribute)
2. [Styleguides](#styleguides)

## How to Contribute

### Bug Reports
Bugs include problems in `master` and `staging` branches. Other branches are considered a work in progress and you can leave a comment on the relevant PR for that branch if there is one open. Please search the open GitHub issues to make sure your bug has not already been reported.

#### What to Include in a Bug Report
- A clear and descriptive title
- The exact steps to reproduce the issue
- Example code or files to reproduce the issue
- A description of the issue and how it differs from the expected behavior

### Feature Suggestions
Feature suggestions include everything from minor improvements to adding completely new functionality to the API. Please search the open GitHub issues to make sure your suggestions has not already been made.

#### What to Include in a Feature Suggestion
- A clear and descriptive title
- An detailed description of the feature
- A description of the API's current behavior and how this feature would change it
- An explanation of why this feature would be useful

### Pull Requests
Take a look at the GitHub issues and leave a comment on one that you would like to work on. We have tagged some issues as `beginner` that are a good introduction to working on the API. Once you have made the appropiate changes open a pull request.

#### What to Include in a Pull Request
- The base branch set to `staging`
- A reference to the issue that this pull request addresses
- A 1 to 3 sentence general summary of the changes made
- Specific changes that were made in checklist format
- Relevent changes to documentation due to code changes

## Styleguides
The following styleguides maintain a consistency across the codebase. Certain aspects are enforced by the continuous build system.

### Go Code Style
- Structurely code should pass `go fmt`
- Methods should named in `CamelCase`
- Exported package variables should be named in `ALL_CAPS`
- Local variables should be named in `snake_case`

### Branch Style
- Feature additions should be prefixed with `feature/`
- Bug fixs should be prefixed with `fix/`
- Refactors should be prefixed with `refactor/`
- Other changes need not be prefixed

### Commit Style
- Each commit should be a complete logical chunk of code
- Squash and rebase commits if needed
- Use past tense in commit messages
- Keep messages clear and concise
- Do not end messages with a period

## Special Thanks
Much of this document was adapted from the [now depreciated HackIllinois API](https://github.com/HackIllinois/api-2018/blob/master/CONTRIBUTING.md).
