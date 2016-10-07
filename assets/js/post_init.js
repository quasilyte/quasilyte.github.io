"use strict";

$(function() {
  function renderPost(post) {
    $("#post-info").html(`
      <div class="pure-g">
        <div class="pure-u-2-3">
          <h2>Post: ${post.title}</h2>
        </div>
        <div class="pure-u-1-3" style="text-align: right">
          <h2>${App.postDateString(post)}</h2>          
        </div>
      </div>
    `);

    $("#footer").html(`
      <br>
      Post tagged with: ${post.tags.join(", ")}
      <hr>      
      <a href="../index.html">Go to main page</a>
    `);
  } 

  var queryParams = window.location.search.substring(1).split("&");
  var postId = ( 
    _.find(queryParams, x => x.startsWith("id"))
    .split("=")[1]
  );

  renderPost(App.posts[postId]);
});