# gh-pr-comments

This is just a hack for myself to avoid context switching between browser and terminal.

## Usage

Pipe the output of

```
gh api -H "Accept: application/vnd.github+json" -H "X-GitHub-Api-Version: 2022-11-28" /repos/OWNER/REPO/pulls/PULL_NUMBER/comments
```

to `gh-pr-comments` and it will pretty-print the comments of the PR in question so you don't need to switch between your browser and terminal when addressing review comments.

```
{file}:{line}
  @{username}: {contents of message}

{file}:{line}
  {username}: {contents of message}
  {username}: {contents of message}
```
