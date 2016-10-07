"use strict";

(function() {
  var globalVars = {
    "Ragf": `<a href="https://github.com/Quasilyte/RAGF">Ragf</a>`,
  };

  var globalsRx = /{(\w+)}/g;
  var localsRx = /{{(\w+)}}/g;
  var linksRx = /\[(\w+):(.*?)\]/g;

  function enrichContent(links, localVars) {
    links = links || {};
    localVars = localVars || {};

    $(".rich").each(function() {
      var html = this.innerHTML;

      // Replace global vars
      html = html.replace(globalsRx, (_, key) => globalVars[key]);
      // Replace links with anchors 
      html = html.replace(linksRx, function(_, key, text) {
        return `<a href="${links[key]}">${text}</a>`;
      });
      // Replace local vars 
      html = html.replace(localsRx, (_, key) => localVars[key]);

      this.innerHTML = html;
    });
  }

  App.enrichContent = enrichContent;
}());