# YouTube Music RPC

Displays the current track from YouTube Music in your Discord status.

## Preview

- Track title
- Artist name
- Album cover
- Listening time
- "Open YouTube Music" button

## Installation

### 1. Download

- Download `ytmusic-rpc.exe` from [Releases](../../releases)
- Install [Tampermonkey](https://www.tampermonkey.net/) extension in your browser

### 2. Install Tampermonkey script

1. Open Tampermonkey → **Create a new script**
2. Paste the script below
3. Save (`Ctrl+S`)

### 3. Run

1. Launch `ytmusic-rpc.exe`
2. Open [music.youtube.com](https://music.youtube.com)
3. Play any track

## Requirements

- Windows
- Discord (must be running)
- Browser with Tampermonkey

## Tampermonkey Script

```js
// ==UserScript==
// @name         YouTube Music to Go RPC
// @namespace    http://tampermonkey.net
// @version      2.2
// @description  Track sender script to localhost:8080
// @author       You
// @match        https://music.youtube.com/*
// @run-at       document-end
// @grant        GM_xmlhttpRequest
// ==/UserScript==

(function() {
    'use strict';

    let lastTrack = "";

    function checkTrack() {
        const titleEl = document.querySelector('.ytmusic-player-bar .title');
        const artistEl = document.querySelector('.ytmusic-player-bar .byline');
        const imgEl = document.querySelector('ytmusic-player-bar img');
        if (titleEl && artistEl) {
            const title = titleEl.textContent.trim();
            const artist = artistEl.textContent.trim();
            const currentTrack = title + " - " + artist;

            if (currentTrack !== lastTrack && title !== "") {
                lastTrack = currentTrack;

                GM_xmlhttpRequest({
                    method: "POST",
                    url: "http://localhost:8080/track",
                    headers: { "Content-Type": "application/json" },
                    data: JSON.stringify({
                        title: title,
                        artist: artist,
                        img: imgEl ? imgEl.src : ""
                    }),
                    onload: function(response) {
                        console.log("Send to Go: " + currentTrack);
                    },
                    onerror: function(err) {
                        console.log("Go server not responding on localhost:8080");
                    }
                });
            }
        }
    }

    setInterval(checkTrack, 2000);
})();
```

## How it works

`ytmusic-rpc.exe` runs a local server on port `8080`. The Tampermonkey script detects the current track on YouTube Music and sends it to the server, which then updates your Discord Rich Presence.




# YouTube Music RPC

Отображает текущий трек из YouTube Music в статусе Discord.

## Как выглядит

- Название трека
- Исполнитель
- Обложка альбома
- Время прослушивания
- Кнопка "Open YouTube Music"

## Установка

### 1. Скачай файлы

- `ytmusic-rpc.exe` — основная программа
- Установи расширение [Tampermonkey](https://www.tampermonkey.net/) в браузер

### 2. Установи скрипт в Tampermonkey

1. Открой Tampermonkey → **Создать новый скрипт**
2. Вставь содержимое файла `ytmusic-rpc.user.js`
3. Сохрани (`Ctrl+S`)

### 3. Запусти

1. Запусти `ytmusic-rpc.exe`
2. Открой [music.youtube.com](https://music.youtube.com)
3. Включи любой трек

## Требования

- Windows
- Discord (должен быть запущен)
- Браузер с Tampermonkey

## Скрипт для Tampermonkey

```js
// ==UserScript==
// @name         YouTube Music to Go RPC
// @namespace    http://tampermonkey.net
// @version      2.2
// @description  Track sender script to localhost:8080
// @author       You
// @match        https://music.youtube.com/*
// @run-at       document-end
// @grant        GM_xmlhttpRequest
// ==/UserScript==

(function() {
    'use strict';

    let lastTrack = "";

    function checkTrack() {
        const titleEl = document.querySelector('.ytmusic-player-bar .title');
        const artistEl = document.querySelector('.ytmusic-player-bar .byline');
        const imgEl = document.querySelector('ytmusic-player-bar img');
        if (titleEl && artistEl) {
            const title = titleEl.textContent.trim();
            const artist = artistEl.textContent.trim();
            const currentTrack = title + " - " + artist;

            if (currentTrack !== lastTrack && title !== "") {
                lastTrack = currentTrack;

                GM_xmlhttpRequest({
                    method: "POST",
                    url: "http://localhost:8080/track",
                    headers: { "Content-Type": "application/json" },
                    data: JSON.stringify({
                        title: title,
                        artist: artist,
                        img: imgEl ? imgEl.src : ""
                    }),
                    onload: function(response) {
                        console.log("Send to Go: " + currentTrack);
                    },
                    onerror: function(err) {
                        console.log("Go server not responding on localhost:8080");
                    }
                });
            }
        }
    }

    setInterval(checkTrack, 2000);
})();
```
