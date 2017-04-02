"use strict";

(function() {
  App.postDateString = function(post) {
    var date = new Date(post.date * 1000);

    var dd = date.getDate();
    var mm = date.getMonth() + 1;
    var yyyy = date.getFullYear();

    dd = dd < 10 ? `0${dd}` : dd;
    mm = mm < 10 ? `0${mm}` : mm;

    return `${dd}.${mm}.${yyyy}`;
  };
}());
