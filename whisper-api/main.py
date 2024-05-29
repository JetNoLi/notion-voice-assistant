from typing import Annotated
import whisper
from fastapi import FastAPI, UploadFile

model = whisper.load_model("small")

app = FastAPI()


@app.get("/")
async def root():
    return "Hello World!"


@app.post("/transcribe/")
async def create_file(file: UploadFile):

    # TODO: Investigate how to process without saving
    # Otherwise use random x chars suffix, to create x distinct temporary files, to allow for x operations
    temp_file = "temp.m4a"
    with open(temp_file, "wb") as f:
        f.write(await file.read())

    # TODO: Look into FP16 not supported warning
    result = model.transcribe(temp_file)

    return {"message": "Completed successfully", "Result": result["text"]}
