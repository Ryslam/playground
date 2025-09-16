import os
import signal
import numpy as np
import pyaudio
from dotenv import load_dotenv
from elevenlabs.client import ElevenLabs
from elevenlabs.conversational_ai.conversation import Conversation
from elevenlabs.conversational_ai.default_audio_interface import DefaultAudioInterface
from openwakeword.model import Model

load_dotenv()

ELEVENLABS_AGENT_ID = os.getenv("ELEVENLABS_AGENT_ID")
ELEVENLABS_API_KEY = os.getenv("ELEVENLABS_API_KEY")

def main():
    oww_model = Model(
        wakeword_models=["hey_mycroft"],
        inference_framework="onnx"
    )

    pa = pyaudio.PyAudio()
    CHUNK = 1280
    RATE = 16000
    CHANNELS = 1
    FORMAT = pyaudio.paInt16

    audio_stream = pa.open(
        format=FORMAT,
        channels=CHANNELS,
        rate=RATE,
        input=True,
        frames_per_buffer=CHUNK
    )

    elevenlabs = ElevenLabs(api_key=ELEVENLABS_API_KEY)

    print("Listening for \"hey_mycroft\"...")

    try:
        while True:
            audio_chunk = np.frombuffer(audio_stream.read(CHUNK), dtype=np.int16)
            prediction = oww_model.predict(audio_chunk)

            for model_name, score in prediction.items():
                if score > 0.5:
                    print(f"Wake word detected: {model_name} (Score: {score:.2f})")
                    oww_model.reset()
                    audio_stream.stop_stream()
                    audio_stream.close()

                    conversation = Conversation(
                        elevenlabs,
                        ELEVENLABS_AGENT_ID,
                        requires_auth=bool(ELEVENLABS_API_KEY),
                        audio_interface=DefaultAudioInterface(),
                        callback_agent_response=lambda response: print(f"Agent: {response}"),
                        callback_agent_response_correction=lambda original, corrected: print(f"Agent: {original} -> {corrected}"),
                        callback_user_transcript=lambda transcript: print(f"User: {transcript}"),
                    )
                    conversation.start_session()

                    signal.signal(signal.SIGINT, lambda sig, frame: conversation.end_session())
                    conversation_id = conversation.wait_for_session_end()
                    print(f"Stopping agent conversational stream. Conversation ID: {conversation_id}")

                    audio_stream = pa.open(
                        format=FORMAT,
                        channels=CHANNELS,
                        rate=RATE,
                        input=True,
                        frames_per_buffer=CHUNK
                    )
                    print("Listening for 'hey_mycroft'...")

    except KeyboardInterrupt:
        print("Stopping wake word detection.")
    finally:
        if audio_stream is not None and not audio_stream.is_stopped():
            audio_stream.stop_stream()
            audio_stream.close()
        if pa is not None:
            pa.terminate()

if __name__ == "__main__":
    main()
