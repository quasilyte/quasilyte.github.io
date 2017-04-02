#!/bin/bash

# 1) concatenate all js/css files into one
cat assets/js/jquery/jquery.min.js \
    assets/js/underscore/underscore.min.js \
    assets/js/init.js \
    assets/js/post.js \
    assets/js/posts_data.js \
    assets/js/post_init.js \
    assets/js/rich_content.js \
    > posts/app.js

cat assets/css/pure/pure-min.css \
    assets/css/style.css \
    > posts/style.css

# 2) create symlincs to resources
ln -sf `pwd`/assets/fonts/* posts/
ln -sf `pwd`/assets/fonts/* assets/css/