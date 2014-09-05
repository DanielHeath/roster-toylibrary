# integREAtor

## Commit notifications

Add a github webhook.

The target URL is `http://integreator-instance.cpdev.realestate.com.au:8080/commit_to_slack/ROOM`

ROOM is the slack room (including a hash encoded as %23) or username (starting with an @) to be notified.

For instance, `%23locations-dev` will post commits to `#locations`,
while  `@daniel-heath` will direct-message me about commits.

## Pull Request notifications

Add a github webhook.

The target URL is `http://integreator-instance.cpdev.realestate.com.au:8080/pull_request_to_slack/ROOM`
