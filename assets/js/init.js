"use strict";

var App = {};

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