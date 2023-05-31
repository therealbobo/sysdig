# Contributing

## Developer Certification of Origin (DCO)

The Apache 2.0 license tells you what rights you have that are provided by the copyright holder. It is important that the contributor fully understands what rights they are licensing and agrees to them. Sometimes the copyright holder isn't the contributor, such as when the contributor is doing work on behalf of a company.

To make a good faith effort to ensure these criteria are met, we require the Developer Certificate of Origin (DCO) process to be followed.

The DCO is an attestation attached to every contribution made by every developer. In the commit message of the contribution, the developer simply adds a Signed-off-by statement and thereby agrees to the DCO, which you can find at http://developercertificate.org.

### DCO Sign-Off Methods

The DCO requires a sign-off message in the following format appear on each commit in the pull request:

```
Signed-off-by: John Doe <john.doe@sysdig.com>
```

You have to use your real name (sorry, no pseudonyms or anonymous contributions).

The DCO text can either be manually added to your commit body, or you can add either `-s` or `--signoff` to your usual `git commit` commands. If you are using the GitHub UI to make a change, you can add the sign-off message directly to the commit message when creating the pull request. If you forget to add the sign-off you can also amend a previous commit with the sign-off by running `git commit --amend -s`. If you've pushed your changes to GitHub already you'll need to force push your branch after this with `git push -f`.

