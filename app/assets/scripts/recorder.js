let audioRecorder = document.getElementById("audio-recorder")
let audioFileInput = document.getElementById("audio-file")
let stream = {}
let mediaRecorder = {}
let audioChunks = [];
let recording = false;

const createRecorder = async (fileElementId) => {
    if (!navigator.mediaDevices.getUserMedia) {
        alert("cannot record audio on this device")
        return
    }

    if (!window.MediaRecorder) {
        alert("media recorder not supported")
        return
    }

    if (!audioRecorder) {
        audioRecorder = document.getElementById("audio-recorder")
    }

    if (mediaRecorder.state === "recording") {
        mediaRecorder.stop()
        console.log("track stopped", mediaRecorder)
        return;
    }

    stream = await navigator.mediaDevices.getUserMedia({
        audio: true
    })

    mediaRecorder = new MediaRecorder(stream)
    mediaRecorder.start()

    audioRecorder.onpause = () => {
        console.log("paused")
    }

    audioRecorder.onloadeddata = (e) => {
        console.log("can listen ", e)
    }

    mediaRecorder.ondataavailable = (e) => {
        console.log("pushing chunks", e.data)
        audioChunks.push(e.data)
    }

    mediaRecorder.onstop = (e) => {
        const blob = new Blob(audioChunks, { type: "audio/webm; codecs=opus" });
        audioChunks = [];

        const audioURL = URL.createObjectURL(blob);
        audioRecorder.src = audioURL;

        if (fileElementId !== "none") {
            // Use a DataTransfer object to simulate a file upload
            const file = new File([blob], "audio.wav", { type: blob.type });
            const dataTransfer = new DataTransfer();
            dataTransfer.items.add(file);

            audioFileInput = document.getElementById(fileElementId || "audio-file")
            audioFileInput.files = dataTransfer.files;

            audioRecorder.controls = true;
        }
    }
}