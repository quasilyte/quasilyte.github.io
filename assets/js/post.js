"use strict";

(function() {
  App.postDateString = function(post) {
    return (
      new Date(post.date * 1000).
      toLocaleString("en-us", {
        "year": "numeric", 
        "month": "2-digit", 
        "day": "2-digit"
      }).
      replace(/(\d+)\/(\d+)\/(\d+)/, "$2.$1.$3")
    );
  };
}());