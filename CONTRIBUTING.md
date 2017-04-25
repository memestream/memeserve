# Contributing
So you have too much time and and want to waste it on this stupid idea?
Splendid!

## Forking

1. Fork the repo (duh)
2. Clone this repo ( `git clone git@github.com:memestream/memeserve` )
3. Remove the existing origin ( `git remote remove origin`)
4. Add your fork as origin ( `git remote add origin git@github.com:you/memeserve` )
3. Add upstream ( `git remote add upstream git@github.com:memestream/memeserve`)

## General
Please make sure your code is properly checked with `gofmt`. We will accept both
the standard and `goimports` style imports.

We do not require strict coverage, but tests are more than welcome.

We will occasionally complain about how terrible your coding skills are (then
cry because we are even worse), so do not take it personally.

To make sure the commit history looks reasonable, we suggest you merge upstream
and then create a new branch and make a PR based on that.
