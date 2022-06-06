# discord-notification-payment

I used database firestore, deployed google cloud function and use cloud schuduler for trigger cronjob

.env
GOOGLE_APPLICATION_CREDENTIALS=PATH_TO_SERVICE_KEY
WEBHOOK_ID=DISCORD_WEBHOOK_ID
WEBHOOK_TOKEN=DISCORD_WEBHOOK_TOKEN
NOTIFY_TEMPLATE_PATH=notify-template.txt

.env.yaml (use for google cloud function deploy)
#use prefix serverless_function_source_code path
GOOGLE_APPLICATION_CREDENTIALS=serverless_function_source_code/PATH_TO_SERVICE_KEY
WEBHOOK_ID=DISCORD_WEBHOOK_ID
WEBHOOK_TOKEN=DISCORD_WEBHOOK_TOKEN
NOTIFY_TEMPLATE_PATH=serverless_function_source_code/notify-template.txt
