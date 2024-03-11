FROM python:3.12.2-bookworm

COPY ./starter-project-python .

COPY ./starter-project-python/requirements/backend.in ./requirements/backend.in
RUN pip install -r ./requirements/backend.in

WORKDIR /app
EXPOSE 8080

CMD uvicorn spaceship.main:app --host=0.0.0.0 --port=8080
