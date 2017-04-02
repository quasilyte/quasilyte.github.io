"use strict";

var App = {
  "rodata": {}
};

$(function() {
  function renderHeader() {
    var $header = $("#header");

    $header.html(`
      <h3>
        <a href="https://quasilyte.github.io/">
          Iskander Sharipov technical blog
        </a>
      </h3>
      <hr>
    `);
  }

  renderHeader();
});
