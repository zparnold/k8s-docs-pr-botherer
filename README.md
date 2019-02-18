# K8s Docs PR Botherer

It bothers people who need to sign the CLA (automagically) in the Kubernetes Website repo

## How to use:

If you like `go`, you'll love this:
```shell
$ go get ./...
$ GITHUB_TOKEN= go run main.go
```

You'll need to fill in the `GITHUB_TOKEN=` with a personal access token you generate. 

## How to get a Github Token:

Follow steps 1-10 here: https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/

For step 7 you'll want to give your token `public_repo` access