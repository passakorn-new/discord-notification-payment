# discord-notification-payment

I used database firestore, deployed google cloud function and use cloud schuduler for trigger cronjob

.env <br />
GOOGLE_APPLICATION_CREDENTIALS=PATH_TO_SERVICE_KEY <br />
WEBHOOK_ID=DISCORD_WEBHOOK_ID <br />
WEBHOOK_TOKEN=DISCORD_WEBHOOK_TOKEN <br />
NOTIFY_TEMPLATE_PATH=notify-template.txt <br /> <br />

.env.yaml (use for google cloud function deploy) <br />
#use prefix serverless_function_source_code path <br />
GOOGLE_APPLICATION_CREDENTIALS=serverless_function_source_code/PATH_TO_SERVICE_KEY <br />
WEBHOOK_ID=DISCORD_WEBHOOK_ID <br />
WEBHOOK_TOKEN=DISCORD_WEBHOOK_TOKEN <br /> 
NOTIFY_TEMPLATE_PATH=serverless_function_source_code/notify-template.txt <br />
