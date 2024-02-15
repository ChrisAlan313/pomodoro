
# Pomodoro

You call pomodoro with an audio file or stream url and a count down will take
over your terminal along with streaming your audio source. Hitting `p` will
pause the timer and hitting `space` will pause the audio stream.

## Organization, from Chatgpt

Here's a suggested layout for your project:

### Main Package (main):

The entry point of your application. Responsible for parsing command line
arguments (such as the audio file/stream URL and the number of minutes for the
countdown). Orchestrates the interaction between the audio playback and the
timer display components.

### Audio Playback Package (audio):

Handles the playback of audio files and streaming audio. Should abstract away
the specifics of dealing with different audio formats and sources. Provides a
simple interface to start, stop, and control audio playback.

### Timer Package (timer):

Manages the countdown timer logic. Could provide functionality to start a
timer, check the remaining time, and trigger an event or callback when the
timer expires. Should be independent of the UI/display logic for the timer.

### UI Package (ui):

Responsible for all terminal-based user interface elements. Displays the
countdown timer in a centered format within the terminal. Could be extended to
include additional UI elements in the future (e.g., playback controls, audio
information display). Interacts with the timer package to update the display as
the countdown progresses.

### Utility Package(s) (util, optional):

Contains helper functions and utilities that might be used across different
parts of the application (e.g., error handling utilities, logging). Helps in
keeping the code DRY and simplifies common tasks across the application.

## Package Responsibilities:

### Main Package:

Initializes and configures the application components. Parses user input and
passes it to the appropriate components.

### Audio Playback Package:

Abstracts the complexities of audio file/stream handling. Offers a simple
play/pause interface to the rest of the application.

### Timer Package:

Provides a robust timer implementation that can be started with a specific
duration. Offers callbacks or channels for timer events (e.g., each second
tick, timer completion).

### UI Package:

Manages terminal output, ensuring the countdown is displayed correctly. Updates
the display based on timer ticks without disrupting the user's terminal.

