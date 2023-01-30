# Astroadd

A go cli for astro.
Create page with auto generated frontmatter.

## Command

### Usage

Make sure you run the command inside the root directory of your site projet.
It will create a file in the src/content path.

```bash
# $ astroadd [path]

$ astroadd posts/my-blog-post.md
```
### output

```markdown
<!-- src/pages/posts/my-blog-post.md -->
---
title : "My blog post"
pubDate: 2023-01-02T16:46:46+01:00
description: ""
tags: []
draft: true
---
```

