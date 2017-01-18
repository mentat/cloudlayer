FROM python:alpine

COPY ./requirements.txt .
COPY ./little_mimic.py .

RUN pip install -r requirements.txt

EXPOSE 8080

ENTRYPOINT [ "python", "./little_mimic.py" ]
