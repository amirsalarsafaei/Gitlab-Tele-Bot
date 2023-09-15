# Gitlab-Tele-Bot
Gitlab Telegram Bot

This a webhook server for Gitlab webhook events. It sends important notification to the telegram chat/group/forum.

![sample](https://github.com/amirsalarsafaei/Gitlab-Tele-Bot/blob/main/Sample%20Merge%20Request%20Message.png)

# How To Use
First, install the latest version using Go's intuitive tool.

```bash
go install github.com/amirsalarsafaei/Gitlab-Tele-Bot@v1.0.2
```
### Config
There is 2 ways to give program config:
- create a `notifier-config.yaml` in home directory of your os.
- using environment variables.


### Example config

yaml config
```yaml
secret: somesecret
telegram:
  token: your-telegram-bot-token
  chat_id: the-chat-id(number)
  thread_id: thread-id(number)
```

env config
```env
SECRET=somesecret
TELEGRAM_TOKEN=your-telegram-bot-token
TELEGRAM_CHAT_ID=the-chat-id(number)
TELEGRAM_THREAD_ID=thread-id(number)
```

### How to setup

- Create a bot in telegram using `bot father` and using `getUpdates` method in [telegram api doc](https://core.telegram.org/bots/api) find out the chat id you want the bot to send notifications to.
- Run this program `Gitlab-Tele-Bot serve 8080`(you can change the port as you please) on your vps.
- Add the IP or URL(if you use nginx to put server behind a domain) to GitLab webhooks

### Current state

In the current state, we only support merged merge request events!
