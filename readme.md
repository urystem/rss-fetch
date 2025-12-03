# 📡 RSSHub

**RSSHub** is a lightweight **CLI-based RSS feed aggregator** written in Go.  
It fetches and parses RSS feeds, stores articles locally, and runs background aggregation with a **worker pool** for high performance.  

With **RSSHub**, you can centralize multiple RSS feeds, keep track of the latest articles, and control the aggregator dynamically — all from your terminal.  

---

## 🚀 Installation

Clone the repository and build the binary:

```sh
git clone https://github.com/bsagat/RSSHub.git
cd RSSHub
go build -o rsshub .
```

---

## Configuration

Create a .env file based on provided .env.example file. 
If u want you can set up your own environment variables such as ports, database URLs, etc.


```sh
cp 
```
---

## 🛠️ CLI Usage

```sh
./rsshub COMMAND [OPTIONS]
```

### Available Commands

| Command          | Description                                                           | Example                                                                 |
| ---------------- | --------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `fetch`          | Start background fetching with worker pool                            | `./rsshub fetch`                                                        |
| `status`          | Show current status of application                            | `./rsshub status`                                                              |
| `add`            | Add a new RSS feed                                                    | `./rsshub add --name "tech-crunch" --url "https://techcrunch.com/feed/" --desc "some description"` |
| `list`           | List available RSS feeds (optionally limit with `--num`)              | `./rsshub list --num 5`                                                 |
| `delete`         | Delete a feed by name                                                 | `./rsshub delete --name "tech-crunch"`                                  |
| `articles`       | Show latest articles from a feed (default: 3, configurable with `--num`) | `./rsshub articles --feed-name "tech-crunch" --num 5`                   |
| `set-interval`   | Change fetch interval dynamically                                     | `./rsshub set-interval 2m`                                              |
| `set-workers`    | Resize worker pool dynamically                                        | `./rsshub set-workers 5`                                                |
| `--help`         | Show help message                                                     | `./rsshub --help`                                                       |

---

## 🔄 Example Workflow

### Open terminal 1:

```sh
# Start the aggregator
$ ./rsshub fetch
$ The background process for fetching feeds has started (interval = 3 minutes, workers = 3)
```

### Open terminal 2:

```sh
$ ./rsshub set-interval 2m
Interval of fetching feeds changed from 5 minutes to 10 minutes

$ ./rsshub set-workers 5
Number of workers changed from 3 to 5
```

List feeds:

```sh
$ ./rsshub list --num 5

# Available RSS Feeds

1. Name: tech-crunch
   URL: https://techcrunch.com/feed/
   Added: 2025-06-10 15:34

2. Name: hacker-news
   URL: https://news.ycombinator.com/rss
   Added: 2025-06-10 15:37

3. Name: bbc-world
   URL: http://feeds.bbci.co.uk/news/world/rss.xml
   Added: 2025-06-11 09:15

4. Name: the-verge
   URL: https://www.theverge.com/rss/index.xml
   Added: 2025-06-12 13:50

5. Name: ars-technica
   URL: http://feeds.arstechnica.com/arstechnica/index
   Added: 2025-06-13 08:25
```

Show latest articles:

```sh
$ ./rsshub articles --feed-name "tech-crunch" --num 5

# Feed: tech-crunch

1. [2025-08-17 21:08:21] GPT-5 is supposed to be nicer now
   https://techcrunch.com/2025/08/17/gpt-5-is-supposed-to-be-nicer-now/

2. [2025-08-17 20:07:35] ‘Stranger Things’ creators may be leaving Netflix
   https://techcrunch.com/2025/08/17/stranger-things-creators-may-be-leaving-netflix/

3. [2025-08-17 16:34:17] Duolingo CEO says controversial AI memo was misunderstood
   https://techcrunch.com/2025/08/17/duolingo-ceo-says-controversial-ai-memo-was-misunderstood/

4. [2025-08-16 20:32:36] Judge says FTC investigation into Media Matters ‘should alarm all Americans’
   https://techcrunch.com/2025/08/16/judge-says-ftc-investigation-into-media-matters-should-alarm-all-americans/

5. [2025-08-16 19:39:40] AI-powered stuffed animals are coming for your kids
   https://techcrunch.com/2025/08/16/ai-powered-stuffed-animals-are-coming-for-your-kids/
```

---

## 📡 Example Feeds

You can try RSSHub with these feeds:

- TechCrunch — `https://techcrunch.com/feed/`
- BBC News — `https://feeds.bbci.co.uk/news/world/rss.xml`
- Ars Technica — `http://feeds.arstechnica.com/arstechnica/index`
- The Verge — `https://www.theverge.com/rss/index.xml`


## Author

This project has been created by:

[Urystem](https://github.com/urystem)