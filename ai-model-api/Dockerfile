# TODO: Try to use CUDA

# from ubuntu:20.04 as base
from python:3.10 as base

WORKDIR /app

COPY main.py .
COPY run.py .
COPY requirements.txt .

RUN echo "starting install"

RUN apt update -y

RUN apt install ffmpeg -y

# RUN apt-get install git-lfs

# RUN git lfs install

RUN apt-get install -y gcc g++ procps

RUN apt-get update && apt-get install -y libhdf5-dev

# clone the repo
# RUN apt install python3.9 -y

# RUN apt-get install python3-pip -y

# RUN apt install git -y

# RUN pip install setuptools-rust


# RUN pip install python-multipart
RUN pip install --upgrade pip setuptools wheel
RUN pip install git+https://github.com/openai/whisper.git 
RUN pip install fastapi
# RUN pip install versioned-hdf5

RUN pip install transformers torch tensorflow 
# RUN pip install nvidia-pyindex nvidia-tensorrt
# RUN pip install 

RUN pip install -r requirements.txt

EXPOSE 8000

# ENTRYPOINT ["fastapi", "run", "main.py", "--bind", "0.0.0.0:8000"]
ENTRYPOINT [ "uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000" ]