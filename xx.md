<details><summary>example 3 (click to expand)</summary>
  
  * Input
    * .wav audio file
    * output directory
  * Generated output description:  
  
  ```text
audio_filename_fft_speech_fe.csv  -- FFT domain 

audio_filename_landmarks.csv  -- C-LMK domain features

audio_filename_landmarks_summary.csv  -- C-LMK summary temporary file

audio_filename_only_vad.csv  -- VAD Onset-Offset durations

audio_filename_pylmk.csv  -- Py-LMK domain features 

audio_filename_pylmk_summary.csv  -- Py-LMK summary temporary file

audio_filename_speech_fe.csv  -- Speech features on whole wav sample 

audio_filename.csv  -- Only Active Voice Speech features based on VAD 
	(audio_filename_speech_fe.csv -> VAD -> audio_filename.csv)

summary.csv  -- Statastical Summary 
	(audio_filename.csv -> stats calculations -> summary.csv)
  ```
  
  `note the newlines and indents - and the \ was added in front of ``` to escape the code-section (remove it for actual use)`
</details>
