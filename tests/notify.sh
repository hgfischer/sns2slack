#!/bin/bash

curl \
	-X POST \
	--data @notify.txt \
	--header 'Content-Type: text/plain; charset=UTF8' \
	--header 'x-amz-sns-message-type: Notification' \
	--header 'x-amz-sns-message-id: da41e39f-ea4d-435a-b922-c6aae3915ebe' \
	--header 'x-amz-sns-topic-arn: arn:aws:sns:us-east-1:123456789012:MyTopic' \
	--header 'x-amz-sns-subscription-arn: arn:aws:sns:us-east-1:123456789012:MyTopic:2bcfbf39-05c3-41de-beaa-fcfcc21c8f55' \
	--header 'User-Agent: Amazon Simple Notification Service Agent' \
	http://127.0.0.1:8080/hook/team/token/tests/:cloud:/sns/publish
