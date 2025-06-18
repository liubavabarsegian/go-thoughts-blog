 todo:
 - [ ] Pagination (show no more than 3 thoughts on one page for people to not get too much of your awesomeness)
 - [ ] Any additional files (images, css, js, if you decide to use any of those) should be submitted as a *zip* file to be unpacked in the same directory as the binary itself, resulting in something like this:
 ```
├── css
│   └── main.css
├── images
│   └── my_cat.png
├── js
│   └── scripts.js
└── myblog-binary
```
- [ ] Admin credentials for posting access (login and password) and database credentials (database name and user) should also be provided separately in a file called *admin_credentials.txt*. If there are additional commands that need to be executed to create tables in a database, put them in the same file.

- [ ] The main page should contain your logo from EX00, links to articles and (optionally) a short preview of their content, as well as pagination (if there are more than 3 articles in a database).

- [ ] When clicking on a link to an article, the user should be taken to a page with rendered markdown text and a single "back" link that takes him/her back to the main page.

- [ ] rate limiter "429 Too Many Requests"
