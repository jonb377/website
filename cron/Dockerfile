FROM python:3.7.4

COPY requirements.txt .

RUN pip3 install -r requirements.txt \
    && rm requirements.txt \
    && apt-get update && apt-get install -y cron \
    && touch /var/log/cron.log

COPY scripts scripts
COPY config/cronjobs /etc/cron.d/cronjobs

RUN crontab /etc/cron.d/cronjobs

CMD ["cron", "-f"]
