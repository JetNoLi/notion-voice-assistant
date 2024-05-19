from typing import Annotated
import whisper
from fastapi import FastAPI, File, UploadFile
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer, pipeline


def use_llama(prompt, max_length):

    access_token="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" #TODO: Env var
            
    # Create the model and tokenizer
    tokenizer = AutoTokenizer.from_pretrained("meta-llama/Llama-2-70b-chat-hf")
    # gpt_model = AutoModelForCausalLM.from_pretrained("meta-llama/Llama-2-70b-chat-hf")

    # Create the pipeline
    text_gen = pipeline(
        "text-generation",
        model="meta-llama/Llama-2-70b-chat-hf",
        token=access_token,
        tokenizer=tokenizer,
        torch_dtype=torch.float16,
        device_map="auto",)
    
    # Run the model
    sequences = text_gen(
        prompt,
        do_sample=True,
        top_k=10,
        num_return_sequences=1,
        eos_token_id=tokenizer.eos_token_id,
        max_length=max_length,
    )
 
    return sequences


print("Starting Up")

model = whisper.load_model("medium")

app = FastAPI()


@app.get("/")
async def root():
    return {"hello world": "yeah"}

    
@app.post("/files/")
async def create_file(file: UploadFile):
    print("starting")
    
    temp_file = "temp.m4a"
    with open( temp_file, "wb") as f:
        f.write(await file.read())

    print("created temp file")
    
    result = model.transcribe(temp_file)
    
    print("Transcribed: " + result["text"])

    understanding = use_llama(result, 250) # TODO: make max length var
    
    print(understanding)
    
    return {"message": "Completed successfully", "Result": result["text"]}


    