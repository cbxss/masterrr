# masterrr
 Use Soundcloud's mastering service for free.
## Installation
Install `ffmpeg`.

1. Go to soundcloud.com/you/likes and press F12 and click on the network tab.
2. Refresh the page, and look for a network request that accesses  https://api-v2.soundcloud.com/
3. Go to Headers -> Request Headers. 
4. Copy what it says after Authorization: , and paste in .env file. 
5. It should look something like AUTH=OAUTH 2-asdf-asdf-asdf

## Usage

1. Use doubler.exe by dragging your track onto it. This just duplicates the track and concats it to the end of itself.
2. Upload this to soundcloud.
3. Go the track and click Master.
2. Copy the number after https://soundcloud.com/you/mastering/tracks/ in the site URL.
3. Run the program and paste the number when promped. 
4. Chose the audio profile to download by typing the number next to it.

Using the doubler.exe is needed as soundcloud's mastering API only allows half of the track to be mastered. 
