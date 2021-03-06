# Alerting

### Slack

Flagger can be configured to send Slack notifications:

```bash
helm upgrade -i flagger flagger/flagger \
--set slack.url=https://hooks.slack.com/services/YOUR/SLACK/WEBHOOK \
--set slack.channel=general \
--set slack.user=flagger
```

Once configured with a Slack incoming **webhook**, Flagger will post messages when a canary deployment 
has been initialised, when a new revision has been detected and if the canary analysis failed or succeeded.

![Slack Notifications](https://raw.githubusercontent.com/weaveworks/flagger/master/docs/screens/slack-canary-notifications.png)

A canary deployment will be rolled back if the progress deadline exceeded or if the analysis reached the 
maximum number of failed checks:

![Slack Notifications](https://raw.githubusercontent.com/weaveworks/flagger/master/docs/screens/slack-canary-failed.png)

### Microsoft Teams

Flagger can be configured to send notifications to Microsoft Teams:

```bash
helm upgrade -i flagger flagger/flagger \
--set msteams.url=https://outlook.office.com/webhook/YOUR/TEAMS/WEBHOOK
```

Flagger will post a message card to MS Teams when a new revision has been detected and if the canary analysis failed or succeeded:

![MS Teams Notifications](https://raw.githubusercontent.com/weaveworks/flagger/master/docs/screens/flagger-ms-teams-notifications.png)

And you'll get a notification on rollback:

![MS Teams Notifications](https://raw.githubusercontent.com/weaveworks/flagger/master/docs/screens/flagger-ms-teams-failed.png)

### Prometheus Alert Manager

Besides Slack, you can use Alertmanager to trigger alerts when a canary deployment failed:

```yaml
  - alert: canary_rollback
    expr: flagger_canary_status > 1
    for: 1m
    labels:
      severity: warning
    annotations:
      summary: "Canary failed"
      description: "Workload {{ $labels.name }} namespace {{ $labels.namespace }}"
```

### Event Webhook

Flagger can be configured to send event payloads to a specified webhook:

```bash
helm upgrade -i flagger flagger/flagger \
--set eventWebhook=https://example.com/flagger-canary-event-webhook
```

The environment variable *EVENT_WEBHOOK_URL* can be used for activating the event-webhook, too.
This is handy for using a secret to store a sensible value that could contain api keys for example.

When configured, every action that Flagger takes during a canary deployment will be sent as JSON via an HTTP POST
request. The JSON payload has the following schema:

```json
{
  "name": "string (canary name)",
  "namespace": "string (canary namespace)",
  "phase": "string (canary phase)",
  "metadata": {
    "eventMessage": "string (canary event message)",
    "eventType": "string (canary event type)",
    "timestamp": "string (unix timestamp ms)"
  }
}
```

Example:

```json
{
  "name": "podinfo",
  "namespace": "default",
  "phase": "Progressing",
  "metadata": {
    "eventMessage": "New revision detected! Scaling up podinfo.default",
    "eventType": "Normal",
    "timestamp": "1578607635167"
  }
}
```

The event webhook can be overwritten at canary level with:

```yaml
  canaryAnalysis:
    webhooks:
      - name: "send to Slack"
        type: event
        url: http://event-recevier.notifications/slack
```
