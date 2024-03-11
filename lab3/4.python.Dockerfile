FROM python:3.12.2-alpine3.19

WORKDIR /app
EXPOSE 8080

COPY ./starter-project-python/requirements/backend.in ./requirements/backend.in
RUN pip install -r ./requirements/backend.in

COPY ./starter-project-python .

CMD uvicorn spaceship.main:app --host=0.0.0.0 --port=8080
