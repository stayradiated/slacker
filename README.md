# slacker

> Post notifications on a slack channel.

You'll need to [setup an incoming
webhook](https://api.slack.com/incoming-webhooks) on Slack first, in order to
get the correct URL.

```
slacker := &Slacker{
  URL: "https://hooks.slack.com/...",
  Icon:     "http://i.imgur.com/C2oopzo.png",
  Username: "Pikachu",
}

slacker.Send("Pika Pika!")
```
