#!/bin/bash

# 1) concatenate all js/css files into one
cat assets/js/jquery/jquery.min.js \
    assets/js/underscore/underscore.min.js \
    assets/js/init.js \
    assets/js/post.js \
    assets/js/posts_data.js \
    assets/js/rich_content.js \
    > posts/app.js

cat assets/css/pure/pure-min.css \
    assets/css/style.css \
    > posts/style.css

# 2) transpile to es5

# 3) minify es5
