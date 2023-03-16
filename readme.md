# Take-home project for Software Developer position @Honeybadger

by Rahmat Hanafi

This is a simple Go program that listens for incoming HTTP POST requests containing a JSON payload, checks if the payload is a spam notification, and sends an alert to a specified Slack channel if it is.

## Installation

- Clone this repository to server
- Install the necessary dependencies by running: `go mod download`

## Usage

1. Set the environment variables in `config.json`
2. Start the program by build with `go build` and run the executable file

## Testing

- Endpoint: `"/"` [POST]
- To test the program, send an HTTP POST request to Endpoint with a JSON payload containing the spam notification. If the program detects a spam notification, it will send an alert to the specified Slack channel.

## Payload format

The program expects incoming HTTP POST requests with a JSON payload containing the following fields:

- `RecordType`: a string representing the type of record (e.g., "Bounce").
- `Type`: a string representing the type of notification (e.g., "SpamNotification").
- `TypeCode`: an integer representing the code for the notification type.
- `Name`: a string representing the name of the notification type (e.g., "Spam notification").
- `Tag`: a string representing a tag associated with the notification.
- `MessageStream`: a string representing the name of the message stream (e.g., "outbound").
- `Description`: a string describing the notification.
- `Email`: a string representing the email address that the notification applies to.
- `From`: a string representing the email address that the notification was sent from.
- `BouncedAt`: a string representing the time at which the notification was bounced.

### Example 1

The following payload should result in an alert being sent to the specified Slack channel:

```json
{
  "RecordType": "Bounce",
  "Type": "SpamNotification",
  "TypeCode": 512,
  "Name": "Spam notification",
  "Tag": "",
  "MessageStream": "outbound",
  "Description": "The message was delivered, but was either blocked by the user, or classified as spam, bulk mail, or had rejected content.",
  "Email": "zaphod@example.com",
  "From": "notifications@honeybadger.io",
  "BouncedAt": "2023-02-27T21:41:30Z"
}
```

### Example 2

The following payload is spam notification but result is not being sent because error on pushing to Slack:

```json
{
  "RecordType": "Bounce",
  "Type": "SpamNotification",
  "TypeCode": 512,
  "Name": "Spam notification",
  "Tag": "",
  "MessageStream": "outbound",
  "Description": "The message was delivered, but was either blocked by the user, or classified as spam, bulk mail, or had rejected content.",
  "Email": "notASlackUser@email.com",
  "From": "notifications@honeybadger.io",
  "BouncedAt": "2023-02-27T21:41:30Z"
}
```

### Example 3

The following payload should not result in an alert being sent:

```json
{
  "RecordType": "Bounce",
  "MessageStream": "outbound",
  "Type": "HardBounce",
  "TypeCode": 1,
  "Name": "Hard bounce",
  "Tag": "Test",
  "Description": "The server was unable to deliver your message (ex: unknown user, mailbox not found).",
  "Email": "arthur@example.com",
  "From": "notifications@honeybadger.io",
  "BouncedAt": "2019-11-05T16:33:54.9070259Z"
}
```
