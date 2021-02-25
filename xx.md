![](./.Sonde_logo.png)

# SP Feature Extraction Library

Supported Features

|                |               |
| -------------- | ------------- |
| mfcc           | jitterddp     |
| lsp            | shimmer       |
| fmt            | spectralslope |
| deltamfcc      | lmk           |
| deltadeltamfcc | pylmk         |
| vadannotation  | rmsamp        |
| scf            | minamp        |
| f0             | maxamp        |
| jitter         | avgamp        |



## Requirements

OS: Linux

Python Version: Python3.6 or Python3.7

## Usage

### Run

Clone this branch into your preferred directory

Open terminal in the same directory and run the below command

`python3    main.py    "wav_file"    "output_dir"`

e.g  `python3    main.py    demo.wav    ~/Desktop/sp_fe_out  `

On successful run, you will get the below log,

```bash
start: sp-fe execution...
done:  sp-fe execution...
task finished in xx.xxxx sec
```

### Configure Features

Open manifest.yaml provided in the package

```yaml
feature_set: 
  - mfcc
  - lsp
  - fmt
  - ...
```
Suppose, you don't want to extract features for lsp.
Remove the lsp value from `feature_set` key.
In that case, your manifest.yaml would look like,

```yaml
feature_set: 
  - mfcc
  - fmt
  - ...
```

Now run the package.

<details>
  <summary>Summary</summary>
  * Input:
    * .wav audio file
    * output directory
 
</details>
	



* Input:
  * .wav audio file
  * output directory

* Generated output description:  

```
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

