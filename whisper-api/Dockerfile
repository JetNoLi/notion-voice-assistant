# TODO: Try to use CUDA
from python:3.10 as base

WORKDIR /app

COPY main.py .
COPY requirements.txt .

RUN apt update -y

RUN apt install ffmpeg -y

RUN pip install git+https://github.com/openai/whisper.git 
RUN pip install fastapi

RUN pip install -r requirements.txt

EXPOSE 8000

ENTRYPOINT [ "uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000" ]