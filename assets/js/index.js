"use strict";

$(function() {
  function renderLatestPosts() {
    var $latestPosts = $("#latest-posts");

    var latestPosts = _.chain(App.posts). 
      filter(x => !x.hidden).
      sort((a, b) => a.date > b.date).
      first(5). 
      value();

    var htmlParts = _.map(latestPosts, function(post) {
      var dateString = App.postDateString(post);

      var url = "#";
      if (post.file) {
        url = `posts/${post.file}.html?id=${post.id}`;
      }

      var tagHtml = (function() {
        var tags = post.tags;
        var tagString = tags.join(', ');
        var spanClass = tags.length ? "hover-me" : ""; 
        return (`
          <span class="${spanClass}" title="${tagString}">
            [${tags.length}]
          </span>
        `);
      }());
        
      return (`
        <div class="pure-g white-box">
          <div class="pure-u-1-4">${post.category}</div>
          <div class="pure-u-2-5">
            <a href="${url}">${post.title}</a>
          </div>
          <div class="pure-u-1-8">${tagHtml}</div>          
          <div class="pure-u-1-8">${dateString}</div>
        </div>
      `);
    });

    $latestPosts.html(htmlParts.join(''));
  }

  renderLatestPosts();
});