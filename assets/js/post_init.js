"use strict";

$(function() {
  function renderPost(post) {
    var chapters = _.map($(".chapter"), function(chapter) {
      var $anchor = $(chapter).children();
      return (`
        <li>
          <a href="#${$anchor.prop("name")}">
            ${$anchor.text()}
          </a>
        </li>
      `);
    }).join("");
    chapters = "<ol>" + chapters + "</ol>";

    $("#post-info").html(`
      <div class="pure-g">
        <div class="pure-u-2-3">
          <h2>Post: ${post.title}</h2>
        </div>
        <div class="pure-u-1-3" style="text-align: right">
          <h2>${App.postDateString(post)}</h2>          
        </div>
      </div>
      
      Category: ${post.category}
      <br>
      Content: ${chapters}
    `);

    $("#footer").html(`
      <br>
      Post tagged with: ${post.tags.join(", ")}
      <hr>      
      <nav>
        <table class="pure-table">
          <tr class="dark-color2"><th>
            Navigation
          </th></tr>
          <tr><td>
            Found a typo? Please, 
            <a href="https://github.com/Quasilyte/quasilyte.github.io/issues">fire an issue</a>!<br>
          </td></tr>
          <tr><td>
            Go to 
            <a href="../index.html">main page</a>      
          </td></tr>
        </table>
      </nav>
    `);
  } 

  var queryParams = window.location.search.substring(1).split("&");
  var postId = ( 
    _.find(queryParams, x => x.startsWith("id"))
    .split("=")[1]
  );

  renderPost(App.posts[postId]);
});