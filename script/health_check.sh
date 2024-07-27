#!/usr/bin/bash

SLACK_WEB_HOOK_URL=$SLACK_WEB_HOOK_URL
ENV=$ENV


send_slack_message() {
  curl -X POST -H 'Content-type: application/json' --data "{\"attachments\": [{\"color\": \"#46567f\",\"blocks\": [{\"type\": \"section\",\"text\": {\"type\": \"mrkdwn\",\"text\": \"$1\"}}]}]}" $SLACK_WEB_HOOK_URL
}


health_check() {
    echo "Health check for \`$1\`"
    result=$(curl -s -o /dev/null -w "%{http_code}" $2)
    
    if [ $result -eq 200 ]; then
        result="UP"
    else
        result="DOWN"
    fi

    send_slack_message "Health check for \`$1\` service: \`$result\`"
}

if [ "$ENV" == "prod" ]; then
  echo "NOT IMPLEMENTED"
fi

if [ "$ENV" == "dev" ]; then
    health_check "cms" "http://localhost:5050/actuator/health"
    health_check "iam" "http://localhost:5051/ping"
    health_check "notify" "http://localhost:8000/health"
fi




