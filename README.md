# SNS to Slack

Simple HTTP server that listens for AWS SNS messages and reposts them to a SlackHD webhook.

## Usage

* Build
* Run
* Setup AWS to notify the URL in your sns2slack server:
	/hook/{team}/{token}/{channel}/{icon_emoji}/{username}/publish
	Ex.: /hook/team/your-token/SNS/:cloud:/aws/publish
