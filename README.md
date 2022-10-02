# Go Scapper

A go web scraping framework using json configuration and other customizations to easily scrape websites.

## Input

```json
{
    "levels": [
        {
            "source": "https://mangapill.com/manga/2/one-piece",
            "label": "chapter",
            "save": {
                "type": "directory",
                "name": "Chapter %name%",
                "path": "/"
            },
            "variables": {
                "name": {
                    "selector": "custom",
                    "value": "chapter_name_parser"
                },
                "url": {
                    "selector": "custom",
                    "value": "chapter_url_selector"
                }
            },
            "sort": {
                "by": "name",
                "order": "asc"
            },
            "levels": [
                {
                    "source": "%parent.url%",
                    "label": "image",
                    "save": {
                        "type": "file",
                        "name": "%counter%.jpg",
                        "content": "%file_content_selector%",
                        "path": "%parent.name%/"
                    }
                }
            ]
        }
    ]
}
```

Crwaling happens via levels instead of going through all the links in the root page. We only need to traverse the required links.

For example, if you want to fetch a chapter in a manga, there would be a single level, which contains all the pages of the chapter. If you want to fetch all the chapters of a single manga, you'd have 2 levels - one for fetching all the chapters and another for fetching all the pages in each chapter.

### How it works?

For each level, do the following:

1. Fetch data from `source`
2.

The `sort` attribute is applied after fetching all the `variables` from the source.

The attribute `save` in each level represents what needs to be stored for each objects in that level. The sub-attributes are self-explanotary.

## FAQ

### Why is `levels` an array in the input format?

One level can contain multiple types of data. For example, say you're fetching multiple mangas from a website. Here, for each manga(root level) you'd need:

- All the chapters(which can be saved as a directory)
- Manga cover which would be an image file

As you can see from the above example, it's possible that for each level, you'd need multiple types of data to be fetched. Hence, we define levels as an array.
