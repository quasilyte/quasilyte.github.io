"use strict";

var App = {
  "rodata": {}
};

$(function() {
  function renderHeader() {
    var $header = $("#header");

    $header.html(`
      <h3>Iskander Sharipov personal blog</h3>
      <hr>
    `);
  }

  renderHeader();
});