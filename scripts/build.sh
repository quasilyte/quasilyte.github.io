#!/bin/bash

# 1) concatenate all js/css files into one
cat src/js/jquery/jquery.min.js \
    src/js/underscore/underscore.min.js \
    src/js/init.js \
    src/js/post.js \
    src/js/posts_data.js \
    src/js/post_init.js \
    src/js/rich_content.js \
    > html/post.js

cat src/js/jquery/jquery.min.js \
    src/js/underscore/underscore.min.js \
    src/js/init.js \
    src/js/post.js \
    src/js/posts_data.js \
    src/js/index.js \
    > html/index.js

cat src/assets/css/pure/pure-min.css \
    src/assets/css/style.css \
    > html/style.css

# 2) copy assets 
cp src/assets/fonts/* html/

# 3) copy posts 
cp src/posts/* html/

# 4) prepare index page 
cp src/index.html html/
