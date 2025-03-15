FROM python:3.11
WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .

# Set a default value for WORKERS in case it's not provided
ENV WORKERS=1

#CMD ["gunicorn", "main:app", "--bind", "0.0.0.0:8001", "--worker-class uvicorn.workers.UvicornWorker"]
CMD ["sh", "-c", "gunicorn main:app --bind 0.0.0.0:8000 --workers ${WORKERS} --worker-class uvicorn.workers.UvicornWorker"]
